package dsl

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/hcl/v2/hclsimple"

	"github.com/jurelou/forensibus/utils/processors"
)

type Config struct {
	ArchivesPasswords []string `hcl:"archives_passwords,optional"`
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
	Name       string            `hcl:"name,label"`
	Sourcetype string            `hcl:"sourcetype,optional"`
	Config     map[string]string `hcl:"config,optional"`
}

type OutputConfig struct {
	Type     string `hcl:"type,label"`
	Address  string `hcl:"address"`
	Username string `hcl:"username"`
	Password string `hcl:"password"`
}

type Step struct {
	CurrentFolder string
	NextArtifact  string
	Name          string
}

type WalkCallback func(interface{}, []Step) []Step

//	func SetDefaults(config *Config) {
//		for _, extract := range config.Pipeline.Extracts {
//			fmt.Println(">>>", extract)
//		}
//		return
//	}
func LoadDSLFile(filePath string) (Config, error) {
	var config Config

	err := hclsimple.DecodeFile(filePath, nil, &config)
	if err != nil {
		// log.Fatalf("Failed to load configuration: %s", err)
		return config, fmt.Errorf("could not decode file %s: %w", filePath, err)
	}

	// SetDefaults(&config)
	return config, nil
}

func walk(item interface{}, in []Step, cb WalkCallback) {
	switch itemType := item.(type) {
	case PipelineConfig:
		out := cb(item, in)
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
		out := cb(item, in)
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
		out := cb(item, in)
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
		fmt.Printf("I don't know about type %T!\n", itemType)
	}
}

func WalkPipeline(item PipelineConfig, in []Step, cb WalkCallback) {
	walk(item, in, cb)
}

func LintPipeline(item PipelineConfig) error {
	// Pass a dummy item to each steps to make sure the pipeline is fully traversed
	dummy := make([]Step, 1)
	dummy = append(dummy, Step{Name: "dummy", CurrentFolder: "DummyCurrent", NextArtifact: "DummyNext"})
	var lastErr error
	WalkPipeline(item, dummy, func(item interface{}, in []Step) []Step {
		switch config := item.(type) {
		case ProcessConfig:
			if _, err := processors.Get(config.Name); err != nil {
				lastErr = fmt.Errorf("processor `%s` does not exists", config.Name)
			}
		case FindConfig:
			for _, pattern := range config.Patterns {
				_, err := regexp.Compile(pattern)
				if err != nil {
					lastErr = fmt.Errorf("invalid find file pattern for %s (%s)", config.Name, pattern)
				}
			}
		case ExtractConfig:

			for _, pattern := range config.Patterns {
				_, err := regexp.Compile(pattern)
				if err != nil {
					lastErr = fmt.Errorf("invalid extract pattern for %s (%s)", config.Name, pattern)
				}
			}
		}
		return dummy
	})
	return lastErr
}

func CountPipelineSteps(item PipelineConfig) int {
	// Pass a dummy item to each steps to make sure the pipeline is fully traversed
	dummy := make([]Step, 1)
	dummy = append(dummy, Step{Name: "dummy", CurrentFolder: "DummyCurrent", NextArtifact: "DummyNext"})
	count := 0
	WalkPipeline(item, dummy, func(item interface{}, in []Step) []Step {
		if _, ok := item.(ProcessConfig); ok {
			count++
		}
		return dummy
	})
	return count
}

func LoadAndLint(filePath string) (Config, error) {
	// Load and lint the given pipeline
	config, err := LoadDSLFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("could not load DSL file %s: %w", filePath, err)
	}
	if err := LintPipeline(config.Pipeline); err != nil {
		return Config{}, fmt.Errorf("error while checking pipeline: %w", err)
	}
	return config, nil
}
