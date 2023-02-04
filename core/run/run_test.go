package run_test

import (
	"testing"

	dsl "github.com/jurelou/forensibus/core"
)

func TestLintValid(t *testing.T) {
	// TotoAa()
	config, err := dsl.LoadDSLFile("../../datasets/pipelines/simple.hcl")
	if err != nil {
		t.Errorf("Error while Loading pipeline %s", err.Error())
		return
	}
	err = dsl.LintPipeline(config.Pipeline)
	if err != nil {
		t.Errorf("Simple pipeline should be valid")
		return
	}
}

func TestLintInvalidProcessor(t *testing.T) {
	// TotoAa()
	config, err := dsl.LoadDSLFile("../../datasets/pipelines/invalid_processor.hcl")
	if err != nil {
		t.Errorf("Error while Loading pipeline %s", err.Error())
		return
	}
	err = dsl.LintPipeline(config.Pipeline)
	if err == nil {
		t.Errorf("Pipeline should be invalid")
		return
	}
	if err.Error() != "processor `ThisProcessorWillNeverExists` does not exists" {
		t.Errorf("Invalid lint error message %s", err.Error())
	}
}

// func TestSimpleWalk(t *testing.T) {
// 	// TotoAa()
// 	config, err := core.LoadDSLFile("../datasets/pipelines/simple.hcl")
// 	if err != nil {
// 		t.Errorf("Error while Loading pipeline %s", err.Error())
// 		return
// 	}
// }
