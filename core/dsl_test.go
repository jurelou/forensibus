package dsl_test

import ( // "fmt"
	// "reflect"
	"strings"
	"testing"

	dsl "github.com/jurelou/forensibus/core"
)

func TestInvalidFile(t *testing.T) {
	_, err := dsl.LoadDSLFile("../datasets/this.does.not.exists")

	if err == nil {
		t.Errorf("FindFiles should fail to find 'this.does.not.exists'")
	}
	if !strings.Contains(err.Error(), "The configuration file ../datasets/this.does.not.exists does not exist") {
		t.Errorf("Invalid error message: %s", err.Error())
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

	if !strings.Contains(err.Error(), "Invalid find file pattern for invalid pattern") {
		t.Errorf("Invalid error message: %s", err.Error())
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

	if !strings.Contains(err.Error(), "Invalid extract pattern for invalid ep") {
		t.Errorf("Invalid error message: %s", err.Error())
	}
}
