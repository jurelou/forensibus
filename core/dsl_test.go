package dsl_test

import (
	"strings"
	"testing"

	dsl "github.com/jurelou/forensibus/core"
)

func TestInvalidFile(t *testing.T) {
	p, err := dsl.LoadDSLFile("../datasets/this.does.not.exists")

	if err == nil {
		t.Errorf("FindFiles should fail to find 'this.does.not.exists'")
	}
	if !strings.Contains(err.Error(), "The configuration file ../datasets/this.does.not.exists does not exist") {
		t.Errorf("Invalid error message: %s", err.Error())
	}

	if dsl.CountPipelineSteps(p.Pipeline) != 0 {
		t.Errorf("Invalid number of steps, expected 0")
	}
}

func TestInvalidFindPattern(t *testing.T) {
	config, err := dsl.LoadDSLFile("../datasets/pipelines/invalid_find_pattern.hcl")
	if err != nil {
		t.Errorf("Could not load invalid_find_pattern.hcl")
		return
	}

	err = dsl.LintPipeline(config.Pipeline)
	if err == nil {
		t.Errorf("invalid_find_pattern.hcl should be invalid: %s", err.Error())
	}

	if !strings.Contains(err.Error(), "invalid find file pattern for invalid pattern") {
		t.Errorf("Invalid error message: %s", err.Error())
	}

	stepsCount := dsl.CountPipelineSteps(config.Pipeline)
	if stepsCount != 0 {
		t.Errorf("Invalid number of steps, expected 0 got %d", stepsCount)
	}
}

func TestInvalidExtractPattern(t *testing.T) {
	config, err := dsl.LoadDSLFile("../datasets/pipelines/invalid_extract_pattern.hcl")
	if err != nil {
		t.Errorf("Could not load invalid_extract_pattern.hcl")
		return
	}

	err = dsl.LintPipeline(config.Pipeline)
	if err == nil {
		t.Errorf("invalid_extract_pattern.hcl should be invalid: %s", err.Error())
	}

	if !strings.Contains(err.Error(), "invalid extract pattern for invalid ep") {
		t.Errorf("Invalid error message: %s", err.Error())
	}
	stepsCount := dsl.CountPipelineSteps(config.Pipeline)
	if stepsCount != 0 {
		t.Errorf("Invalid number of steps, expected 0 got %d", stepsCount)
	}
}

func TestLoadAndLintInvalidFile(t *testing.T) {
	config, err := dsl.LoadAndLint("/tmp/this/does/not_exists.hcl")
	if err == nil {
		t.Errorf("Pipeline /tmp/this/does/not_exists.hcl should be invalid")
		return
	}
	stepsCount := dsl.CountPipelineSteps(config.Pipeline)
	if stepsCount != 0 {
		t.Errorf("Invalid number of steps, expected 0 got %d", stepsCount)
	}
}

func TestLoadAndLintInvalidExtractPattern(t *testing.T) {
	config, err := dsl.LoadAndLint("../datasets/pipelines/invalid_extract_pattern.hcl")
	if err == nil {
		t.Errorf("Pipeline invalid_extract_pattern.hcl should be invalid")
		return
	}
	stepsCount := dsl.CountPipelineSteps(config.Pipeline)
	if stepsCount != 0 {
		t.Errorf("Invalid number of steps, expected 0 got %d", stepsCount)
	}
}
