package core

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/decompress"
)

var (
	serverAddress = "localhost:50051"
)

type JobChannels struct {
	// List of input / output channels
	Jobs		chan Job
	JobResults	chan JobResult
}

func find(steps []Step, config FindConfig) []Step {
	var out []Step
	// var latestErr error = nil

	for _, step := range steps {
		files, err := utils.FindFiles(utils.FindFilesParams{Path: step.NextArtifact, PathPatterns: config.Patterns, FileFormats: config.MimeTypes})
		// latestErr = err
		if err != nil {
			utils.Log.Warnf("Error while searching for `%s` files from %s: %w", config.Name, step.NextArtifact, err)
			continue
		}
		filesLen := len(files)
		if filesLen == 0 {
			utils.Log.Warnf("Found 0 `%s` files from %s", config.Name, step.NextArtifact)
		} else {
			utils.Log.Infof("Found %d `%s` files from %s", filesLen, config.Name, step.NextArtifact)
		}

		for _, file := range files {
			out = append(out, Step{NextArtifact: file, CurrentFolder: step.CurrentFolder})

		}
	}
	return out
}

// Finds and extract files matching patterns
func extract(steps []Step, config ExtractConfig) []Step {
	// err :=
	var outs []Step
	files := find(steps, FindConfig{Name: config.Name, Patterns: config.Patterns, MimeTypes: config.MimeTypes})

	for _, in := range files {
		out, err := decompress.Decompress(in.NextArtifact, in.CurrentFolder)
		if err != nil {
			utils.Log.Warnf("Error while decompressing %s: %w", in.NextArtifact, err)
			continue
		}
		outs = append(outs, Step{NextArtifact: out, CurrentFolder: out})

		outLen := len(out)
		if outLen == 0 {
			utils.Log.Warnf("Extracted 0 `%s` files from %s", config.Name, in.NextArtifact)

		} else {
			utils.Log.Infof("Extracted %d `%s` files", outLen, config.Name, in.NextArtifact)
		}

	}
	return outs
}

func process(steps []Step, config ProcessConfig, jobs chan Job) {
	// RunProcessor(steps, config)
	for _, in := range steps {
		fmt.Println("Launch ", in.NextArtifact)
		jobs <- Job{Step: in, Config: config.Config, Name: config.Name}
		fmt.Println("..")
	}
}


func makeInputs(sources []string) ([]Step, error) {
	ins := make([]Step, 0, len(sources))
	for _, source := range sources {
		absPath, err := utils.ConvertRelativePath(source)
		if err != nil {
			return ins, err
		}

		filename := filepath.Base(absPath)
		filenameWhithoutSuffix := strings.TrimSuffix(filename, filepath.Ext(filename))

		volumeName := filepath.VolumeName(absPath)
		absPath = strings.TrimPrefix(absPath, volumeName)

		outputFolder := filepath.Join(utils.Config.OutputFolder, filepath.Dir(absPath), filenameWhithoutSuffix)

		utils.Log.Infof("Writing output file `%s` FROM `%s`", outputFolder, absPath)
		ins = append(ins, Step{CurrentFolder: outputFolder, NextArtifact: absPath})

	}
	return ins, nil
}

func RunPipeline(pipeline PipelineConfig, steps []Step, chans JobChannels) {

	WalkPipeline(pipeline, steps, func(item interface{}, in []Step) ([]Step, error) {

		switch item.(type) {
		case FindConfig:
			findConfig := item.(FindConfig)
			out := find(in, findConfig)
			return out, nil

		case ExtractConfig:
			extractConfig := item.(ExtractConfig)
			out := extract(in, extractConfig)
			return out, nil

		case ProcessConfig:
			processConfig := item.(ProcessConfig)
			process(in, processConfig, chans.Jobs)
			return []Step{}, nil

		}
		return in, nil

	})
}

func Run(pipelineconfigFile string, paths []string) {
	utils.Log.Info("Starting client")

	config, err := LoadDSLFile(pipelineconfigFile)
	if err != nil {
		utils.Log.Warnf(err.Error())
		return
	}
	if err := LintPipeline(config.Pipeline); err != nil {
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

	inputs, err := makeInputs(paths)
	if err != nil {
		utils.Log.Warnf(err.Error())
		return
	}

	// Queue size should be > than workers capacity so that workers are never starving
	queueSize := workers.Capacity()
	chans := JobChannels{Jobs: make(chan Job, queueSize * 2), JobResults: make(chan JobResult, queueSize * 4)}

	workers.Work(chans)
	utils.Log.Infof("Launched a pool of %d workers, total capacity is %d", workers.Size(), workers.Capacity())

	RunPipeline(config.Pipeline, inputs, chans)

	// Pipeline is done, close job channel
	close(chans.Jobs)
	// Wait for last jobs to finish
	workers.Wait()
	// Finally, close results channel
	close(chans.JobResults)
	utils.Log.Infof("Terminated")

	// for r := range chans.JobResults {
	// 	fmt.Println("Got result>>", r)
	// }
	// fmt.Println("--", inputs)
	// execPipeline(config, inputs)

}
