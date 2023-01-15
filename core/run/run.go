package run

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"github.com/pterm/pterm"

	dsl "github.com/jurelou/forensibus/core"
	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/decompress"
)

var serverAddress = "localhost:50051"

type JobChannels struct {
	// List of input / output channels
	Jobs       chan Job
	JobResults chan JobResult
}

func find(steps []dsl.Step, config dsl.FindConfig) []dsl.Step {
	var out []dsl.Step
	// var latestErr error = nil

	for _, step := range steps {
		files, err := utils.FindFiles(utils.FindFilesParams{Path: step.NextArtifact, PathPatterns: config.Patterns, FileFormats: config.MimeTypes})
		// latestErr = err
		if err != nil {
			pterm.Error.Printfln("Error while searching for `%s` files from %s: %w", config.Name, step.NextArtifact, err)
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
		out, err := decompress.Decompress(in.NextArtifact, in.CurrentFolder)
		if err != nil {
			pterm.Error.Printfln("Error while decompressing %s: %w", in.NextArtifact, err)
			continue
		}
		outs = append(outs, dsl.Step{Name: "from_extract", NextArtifact: out, CurrentFolder: out})

		outLen := len(out)
		if outLen == 0 {
			pterm.Warning.Printfln("Extracted 0 `%s` files from %s", config.Name, in.NextArtifact)
		} else {
			pterm.Success.Printfln("Extracted %d `%s` files", outLen, config.Name, in.NextArtifact)
		}

	}
	return outs
}

func process(steps []dsl.Step, config dsl.ProcessConfig, jobs chan Job) {
	// RunProcessor(steps, config)
	for _, in := range steps {
		pterm.Info.Printfln("Running %s processor against %s", config.Name, in.NextArtifact)
		jobs <- Job{Step: in, Config: config.Config, Name: config.Name}
	}
}

func MakeInputs(sources []string) ([]dsl.Step, error) {
	ins := make([]dsl.Step, 0, len(sources))
	for _, source := range sources {
		// Get absolute path
		absPath, err := utils.ConvertRelativePath(source)
		if err != nil {
			return ins, fmt.Errorf("Could not convert %s to relative path: %w", source, err)
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

func RunPipeline(pipeline dsl.PipelineConfig, steps []dsl.Step, chans JobChannels) {
	dsl.WalkPipeline(pipeline, steps, func(item interface{}, in []dsl.Step) []dsl.Step {
		switch item.(type) {
		case dsl.FindConfig:
			findConfig := item.(dsl.FindConfig)
			out := find(in, findConfig)
			return out

		case dsl.ExtractConfig:
			extractConfig := item.(dsl.ExtractConfig)
			out := extract(in, extractConfig)
			return out

		case dsl.ProcessConfig:
			processConfig := item.(dsl.ProcessConfig)
			process(in, processConfig, chans.Jobs)
			return []dsl.Step{}

		}
		return in
	})
}

func MonitorResults(stepsCount int, results <-chan JobResult, finish chan<- bool) {
	globalBar, _ := pterm.DefaultProgressbar.WithTotal(stepsCount).WithTitle("Total").WithRemoveWhenDone(false).Start()

	for result := range results {
		globalBar.Increment()
		if utils.IsErrored(result.Status) {
			utils.Log.Warnf("Error !!!!", result)
			pterm.Error.Printfln("Error while running `%s` against `%s`: %s", result.Job.Name, result.Job.Step.NextArtifact, result.Error)

		} else {
			pterm.Success.Printfln("Successfully ran `%s` against `%s`", result.Job.Name, result.Job.Step.NextArtifact)
		}
	}
	finish <- true
}

func loadDSL(filePath string) (dsl.Config, error) {
	// Load and lint the given pipeline
	config, err := dsl.LoadDSLFile(filePath)
	if err != nil {
		return dsl.Config{}, fmt.Errorf("Could not load DSL file %s: %w", filePath, err)
	}
	if err := dsl.LintPipeline(config.Pipeline); err != nil {
		return dsl.Config{}, fmt.Errorf("Error while checking pipeline: %w", err)
	}
	return config, nil
}

func onSigint(sigint <-chan os.Signal) {
	exit := false
	for range sigint {
		if exit == true {
			os.Exit(1)
		}
		fmt.Println("Caught SIGINT. Press ctrl-C once more to exit")
		exit = true
	}
}

func Run(pipelineconfigFile string, paths []string) {
	config, err := dsl.LoadDSLFile(pipelineconfigFile)
	if err != nil {
		utils.Log.Warnf(err.Error())
		return
	}
	stepsCount := dsl.CountPipelineSteps(config.Pipeline)
	inputs, err := MakeInputs(paths)
	if err != nil {
		utils.Log.Warnf(err.Error())
		return
	}

	workers := new(WorkerPool)
	worker, err := workers.Connect(serverAddress)
	if err != nil {
		utils.Log.Warnf(err.Error())
		return
	}

	utils.Log.Infof("Worker %s (%s) is up! (cap: %d)", worker.Address, worker.Name, worker.Capacity)

	// Queue size should be > than workers capacity so that workers are never starving
	queueSize := workers.Capacity()
	chans := JobChannels{Jobs: make(chan Job, queueSize*2), JobResults: make(chan JobResult, queueSize*2)}

	workers.Work(chans)
	utils.Log.Infof("Launched a pool of %d workers, total capacity is %d", workers.Size(), workers.Capacity())

	finishMonitoring := make(chan bool)
	go MonitorResults(stepsCount, chans.JobResults, finishMonitoring)

	// Catch SIGINT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go onSigint(c)

	RunPipeline(config.Pipeline, inputs, chans)

	// Pipeline is done, close job channel
	close(chans.Jobs)
	// Wait for last jobs to finish
	workers.Wait()
	// Finally, close results channel
	close(chans.JobResults)
	// Wait for monitoring to finish
	<-finishMonitoring

	fmt.Println("\nDone!")
}
