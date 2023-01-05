package core

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Config struct {
	ArchivesPasswords []string       `hcl:"archives_passwords"`
	TemporaryFolder   string         `hcl:"temporary_folder"`
	Pipeline          PipelineConfig `hcl:"pipeline,block"`
	Output            OutputConfig   `hcl:"output,block"`
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

func LoadDSLFile(filePath string) (Config, error) {
	var config Config

	err := hclsimple.DecodeFile(filePath, nil, &config)
	if err != nil {
		// log.Fatalf("Failed to load configuration: %s", err)
		return config, err
	}
	return config, nil

}
