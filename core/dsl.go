package core

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"

	"github.com/jurelou/forensibus/utils/processors"
)

type Config struct {
	// ArchivesPasswords []string       `hcl:"archives_passwords"`
	// TemporaryFolder   string         `hcl:"temporary_folder"`
	Pipeline PipelineConfig `hcl:"pipeline,block"`
	Output   OutputConfig   `hcl:"output,block"`
}

type PipelineConfig struct {
	Name string `hcl:"name,label"`

	Extracts  []ExtractConfig `hcl:"extract,block"`
	Finds     []FindConfig    `hcl:"find,block"`
	Processes []ProcessConfig `hcl:"process,block"`
}

type ExtractConfig struct {
	Name      string   `hcl:"name,label"`
	Patterns  []string `hcl:"patterns"`
	MimeTypes []string `hcl:"mime_types,optional"`

	Extracts  []ExtractConfig `hcl:"extract,block"`
	Finds     []FindConfig    `hcl:"find,block"`
	Processes []ProcessConfig `hcl:"process,block"`
}

type FindConfig struct {
	Name      string   `hcl:"name,label"`
	Patterns  []string `hcl:"patterns"`
	MimeTypes []string `hcl:"mime_types,optional"`

	// FileAllowed bool `hcl:"file_allowed,optional"`
	// DirAllowed bool `hcl:"file_allowed,optional"`

	Extracts  []ExtractConfig `hcl:"extract,block"`
	Finds     []FindConfig    `hcl:"find,block"`
	Processes []ProcessConfig `hcl:"process,block"`
}

type ProcessConfig struct {
	Name   string            `hcl:"name,label"`
	Config map[string]string `hcl:"config,optional"`
}

type OutputConfig struct {
	Type     string `hcl:"type,label"`
	Address  string `hcl:"address"`
	Username string `hcl:"username"`
	Password string `hcl:"password"`
}

type Input struct {
	Current string
	Next    string
}

type WalkCallback func(interface{}, []Input) ([]Input, error)

func LoadDSLFile(filePath string) (Config, error) {
	var config Config

	err := hclsimple.DecodeFile(filePath, nil, &config)
	if err != nil {
		// log.Fatalf("Failed to load configuration: %s", err)
		return config, err
	}
	return config, nil

}

func walk(item interface{}, in []Input, cb WalkCallback) {

	switch t := item.(type) {
	case PipelineConfig:
		out, _ := cb(item, in)
		if len(out) == 0 {
			return
		}
		pipelineConfig := item.(PipelineConfig)

		for _, nestedFinds := range pipelineConfig.Finds {
			walk(nestedFinds, out, cb)
		}
		for _, nestedProcess := range pipelineConfig.Processes {
			walk(nestedProcess, out, cb)
		}
		for _, nestedExtracts := range pipelineConfig.Extracts {
			walk(nestedExtracts, out, cb)
		}

	case FindConfig:
		out, _ := cb(item, in)
		if len(out) == 0 {
			return
		}
		findConfig := item.(FindConfig)
		for _, nestedFinds := range findConfig.Finds {
			walk(nestedFinds, out, cb)
		}
		for _, nestedProcess := range findConfig.Processes {
			walk(nestedProcess, out, cb)
		}
		for _, nestedExtracts := range findConfig.Extracts {
			walk(nestedExtracts, out, cb)
		}
	case ExtractConfig:
		out, _ := cb(item, in)
		if len(out) == 0 {
			return
		}
		extractConfig := item.(ExtractConfig)

		for _, nestedFinds := range extractConfig.Finds {
			walk(nestedFinds, out, cb)
		}
		for _, nestedProcess := range extractConfig.Processes {
			walk(nestedProcess, out, cb)
		}
		for _, nestedExtracts := range extractConfig.Extracts {
			walk(nestedExtracts, out, cb)
		}
	case ProcessConfig:
		cb(item, in)
	default:
		fmt.Printf("I don't know about type %T!\n", t)
	}
}

func WalkPipeline(item PipelineConfig, in []Input, cb WalkCallback) {
	walk(item, in, cb)
}

func LintPipeline(item PipelineConfig) error {
	// Pass a dummy item to each steps to make sure the pipeline is fully traversed
	dummy := make([]Input, 1)
	dummy = append(dummy, Input{Current: "DummyCurrent", Next: "DummyNext"})
	var lastErr error

	WalkPipeline(item, dummy, func(item interface{}, in []Input) ([]Input, error) {
		switch item.(type) {
		case ProcessConfig:
			processConfig := item.(ProcessConfig)
			if _, err := processors.Get(processConfig.Name); err != nil {
				lastErr = fmt.Errorf("Invalid pipeline definition, processor %s is not found", processConfig.Name)
				return dummy, nil
			}
		}
		return dummy, nil
	})
	return lastErr
}
