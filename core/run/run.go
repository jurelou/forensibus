package run

import (
	"fmt"
	// "os"
	// "os/signal"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/pterm/pterm"

	dsl "github.com/jurelou/forensibus/core"
	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/decompress"
	"github.com/jurelou/forensibus/utils/names_generator"
	"github.com/jurelou/forensibus/worker"
)

var serverAddress = "localhost:50051"

type ProcessStarted struct {
	Name            string
	ProcessId       string
	TasksCount      int
	TasksTerminated int
}

type TaskStarted struct {
	ProcessId     string
	Tag           string
	Step          dsl.Step
	ProcessConfig dsl.ProcessConfig
}

type TaskEnded struct {
	Name      string
	Source    string
	ProcessId string
	Status    uint32
	Errors    []string
}

func find(steps []dsl.Step, config dsl.FindConfig) []dsl.Step {
	var out []dsl.Step
	// var latestErr error = nil

	for _, step := range steps {
		files, err := utils.FindFiles(utils.FindFilesParams{Path: step.NextArtifact, PathPatterns: config.Patterns, FileFormats: config.MimeTypes})
		// latestErr = err
		if err != nil {
			pterm.Error.Printfln("Error while searching for `%s` files from %s: %s", config.Name, step.NextArtifact, err.Error())
			continue
		}
		filesLen := len(files)
		if filesLen == 0 {
			pterm.Warning.Printfln("Found 0 `%s` files from %s", config.Name, step.NextArtifact)
		} else {
			pterm.Success.Printfln("Found %d `%s` files from %s", filesLen, config.Name, step.NextArtifact)
		}

		for _, file := range files {
			out = append(out, dsl.Step{Name: "from_find", NextArtifact: file, CurrentFolder: step.CurrentFolder})
		}
	}
	return out
}

// Finds and extract files matching patterns
func extract(steps []dsl.Step, config dsl.ExtractConfig) []dsl.Step {
	// err :=
	var outs []dsl.Step
	files := find(steps, dsl.FindConfig{Name: config.Name, Patterns: config.Patterns, MimeTypes: config.MimeTypes})

	for _, in := range files {
		inputFilenameWoExt := strings.TrimSuffix(in.NextArtifact, filepath.Ext(in.NextArtifact))

		out, err := decompress.Decompress(in.NextArtifact, filepath.Join(in.CurrentFolder, filepath.Base(inputFilenameWoExt)))
		if err != nil {
			pterm.Error.Printfln("Error while decompressing %s: %s", in.NextArtifact, err.Error())
			continue
		}
		outs = append(outs, dsl.Step{Name: "from_extract", NextArtifact: out, CurrentFolder: out})

		outLen := len(out)
		if outLen == 0 {
			pterm.Warning.Printfln("Extracted 0 `%s` files from %s", config.Name, in.NextArtifact)
		} else {
			pterm.Success.Printfln("Extracted %d `%s` files from %s", outLen, config.Name, in.NextArtifact)
		}

	}
	return outs
}

func RunPipeline(pipeline dsl.PipelineConfig, steps []dsl.Step, tag string, sProcesses chan<- ProcessStarted, sTasks chan<- TaskStarted) {
	dsl.WalkPipeline(pipeline, steps, func(item interface{}, in []dsl.Step) []dsl.Step {
		switch config := item.(type) {
		case dsl.FindConfig:
			out := find(in, config)
			return out

		case dsl.ExtractConfig:
			out := extract(in, config)
			return out

		case dsl.ProcessConfig:
			id := uuid.NewString()

			sProcesses <- ProcessStarted{
				Name:            config.Name,
				ProcessId:       id,
				TasksCount:      len(in),
				TasksTerminated: 0,
			}

			pterm.Info.Printfln("Running %s processor against %d files", config.Name, len(in))
			for _, i := range in {
				sTasks <- TaskStarted{
					ProcessId:     id,
					Tag:           tag,
					Step:          i,
					ProcessConfig: config,
				}
			}
			return []dsl.Step{}
		}
		return in
	})
}

// func onSigint(sigint <-chan os.Signal) {
// 	exit := false

// 	for range sigint {
// 		os.Exit(1)

// 		if exit == true {
// 			os.Exit(1)
// 		}
// 		fmt.Println("Caught SIGINT. Press ctrl-C once more to exit")
// 		exit = true
// 	}
// }


func MakeInputs(sources []string) ([]dsl.Step, error) {
	ins := make([]dsl.Step, 0, len(sources))
	for _, source := range sources {
		// Get absolute path
		absPath, err := utils.ConvertRelativePath(source)
		if err != nil {
			return ins, fmt.Errorf("could not convert %s to relative path: %w", source, err)
		}
		// Get file or folder name
		filename := filepath.Base(absPath)
		// Get file or folder whithout extension
		filenameWhithoutSuffix := strings.TrimSuffix(filename, filepath.Ext(filename))
		// Get file of folder whithout volume name (for windows compatibility)
		absPath = strings.TrimPrefix(absPath, filepath.VolumeName(absPath))

		// Output folder is a concatenation of the configured temp folder + the given path
		outputFolder := filepath.Join(utils.Config.OutputFolder, filepath.Dir(absPath), filenameWhithoutSuffix)

		// utils.Log.Infof("Writing output file `%s` FROM `%s`", outputFolder, absPath)
		ins = append(ins, dsl.Step{Name: "init", CurrentFolder: outputFolder, NextArtifact: absPath})

	}
	return ins, nil
}

func MakeWorkers(disableLocalWorker bool, tStart <-chan TaskStarted, tEnd chan<- TaskEnded) (*WorkerPool, error) {
	workers := new(WorkerPool)

	if !disableLocalWorker {
		port := utils.FindOpenTcpPort()
		if port == 0 {
			return nil, fmt.Errorf("could not find a port to listen on")
		}

		go func() {
			err := worker.RunWorkerServer(port)
			if err != nil {
				pterm.Error.Printfln("Could not run worker: %s", err.Error())
			}
		}()

		lWorker, err := workers.Connect(fmt.Sprintf("localhost:%d", port))
		if err != nil {
			return nil, fmt.Errorf("could not start local worker: %s", err.Error())
		}
		pterm.Info.Printfln("Local worker is up (capacity of %d)", lWorker.Capacity)

	} else {
		pterm.Warning.Printfln("Disabled local worker")
	}

	_, err := workers.Connect(serverAddress)
	if err != nil {
		pterm.Warning.Println(err.Error())
		// return
	}

	workers.Work(tStart, tEnd)
	return workers, nil
}

func Run(pipelineconfigFile string, paths []string, tag string, disableLocalWorker bool) {
	// Configure logger, if local worker is disabled, do not print to stdout
	err := utils.ConfigureLogger(disableLocalWorker)
	if err != nil {
		pterm.Error.Printfln(err.Error())
		return
	}

	if tag == "" {
		tag = names_generator.GetRandomName()
	}

	config, err := dsl.LoadAndLint(pipelineconfigFile)
	if err != nil {
		pterm.Error.Printfln(err.Error())
		return
	}

	stepsCount := dsl.CountPipelineSteps(config.Pipeline)
	inputs, err := MakeInputs(paths)
	if err != nil {
		utils.Log.Warnf(err.Error())
		return
	}

	startedProcesses := make(chan ProcessStarted, stepsCount)
	startedTasks := make(chan TaskStarted, 512)
	endedTasks := make(chan TaskEnded, 512)

	workers, err := MakeWorkers(disableLocalWorker, startedTasks, endedTasks)
	if err != nil {
		pterm.Error.Printfln(err.Error())
		return
	}
	if workers.Capacity() == 0 {
		pterm.Error.Println("No workers available, you can either remove -d flag or start a new worker")
		return
	} else {
		pterm.Success.Printfln("Launched a pool of %d workers, total capacity is %d", workers.Size(), workers.Capacity())
	}
	pterm.Info.Printfln("Running pipeline `%s` (%d processors)", config.Pipeline.Name, stepsCount)
	pterm.Info.Printfln("Using tag `%s`", tag)
	pterm.Info.Printfln("Using splunk index `%s`", utils.Config.Splunk.Index)

	finishMonitoring := make(chan bool)
	go MonitorResults(finishMonitoring, startedProcesses, endedTasks)

	// Catch SIGINT
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// go onSigint(c)

	RunPipeline(config.Pipeline, inputs, tag, startedProcesses, startedTasks)

	// Pipeline is done, close tasks channel
	close(startedProcesses)
	close(startedTasks)
	// Wait for last jobs to finish
	workers.Wait()
	// Finally, close results channel
	close(endedTasks)
	// Wait for monitoring thread to finish
	<-finishMonitoring
	pterm.Info.Println("Done!")
}
