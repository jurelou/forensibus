package core_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/jurelou/forensibus/core"
)

func TestInvalidFile(t *testing.T) {
	_, err := core.LoadDSLFile("../datasets/this.does.not.exists")

	if err == nil {
		t.Errorf("FindFiles should fail to find 'this.does.not.exists'")
	}
	if !strings.Contains(err.Error(), "The configuration file ../datasets/this.does.not.exists does not exist") {
		t.Errorf("Invalid error message: %s", err.Error())
	}
}

func doSomething(i int) int {
	return i + 1
}
func TestSimple(t *testing.T) {
	config, err := core.LoadDSLFile("../datasets/pipelines/nested.hcl")
	if err != nil {
		t.Errorf("Hcl file `simple.hcl` should not raise errors: %s", err.Error())
	}

	if !reflect.DeepEqual(config.Output, core.OutputConfig{Type: "splunk", Address: "localhost", Username: "admin", Password: "admin"}) {
		t.Errorf("`Output` block does not match")
	}

	pipeline := config.Pipeline

	if pipeline.Name != "test_simple_pipeline" {
		t.Errorf("Pipeline name is invalid: %s", pipeline.Name)

	}
	var walkConfig func(item interface{}, i int)
	source := 0

	walkConfig = func(item interface{}, i int) {
		switch t := item.(type) {
		case core.PipelineConfig:
			pipelineConfig := item.(core.PipelineConfig)

			for _, nestedFinds := range pipelineConfig.Finds {
				walkConfig(nestedFinds, i)
			}
			for _, nestedProcess := range pipelineConfig.Processes {
				walkConfig(nestedProcess, i)
			}
			for _, nestedExtracts := range pipelineConfig.Extracts {
				walkConfig(nestedExtracts, i)
			}

		case core.FindConfig:
			findConfig := item.(core.FindConfig)
			fmt.Println(i, "find ", findConfig.Name)
			i = doSomething(i)

			for _, nestedFinds := range findConfig.Finds {
				walkConfig(nestedFinds, i)
			}
			for _, nestedProcess := range findConfig.Processes {
				walkConfig(nestedProcess, i)
			}
			for _, nestedExtracts := range findConfig.Extracts {
				walkConfig(nestedExtracts, i)
			}
		case core.ExtractConfig:
			extractConfig := item.(core.ExtractConfig)
			fmt.Println(i, "find ", extractConfig.Name)
			i = doSomething(i)

			for _, nestedFinds := range extractConfig.Finds {
				walkConfig(nestedFinds, i)
			}
			for _, nestedProcess := range extractConfig.Processes {
				walkConfig(nestedProcess, i)
			}
			for _, nestedExtracts := range extractConfig.Extracts {
				walkConfig(nestedExtracts, i)
			}
		case core.ProcessConfig:
			//i = doSomething(i)
			processConfig := item.(core.ProcessConfig)
			fmt.Println(i, "process ", processConfig.Name)
		default:
			fmt.Printf("I don't know about type %T!\n", t)
		}

	}

	walkConfig(pipeline, source)
	fmt.Println(source)
}
