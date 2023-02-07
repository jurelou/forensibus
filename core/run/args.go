package run

import (
	"fmt"

	"github.com/jurelou/forensibus/utils"
)

type Splunkargs struct {
	Index   string
	Address string
	Token   string
}

type Runargs struct {
	Targets            []string
	PipelineFile       string
	Tag                string
	DisableLocalWorker bool
	Verbose            bool

	Splunk Splunkargs
}

func (r Runargs) Validate() error {
	if !utils.FileExists(r.PipelineFile) {
		return fmt.Errorf("the given pipeline file `%s` does not exists", r.PipelineFile)
	}
	if len(r.Targets) < 1 {
		return fmt.Errorf("requires at least one target filepath argument")
	}
	for _, filepath := range r.Targets {
		if !utils.FileExists(filepath) {
			return fmt.Errorf("target filepath %s does not exists", filepath)
		}
	}
	return nil
}

func NewRunargs() *Runargs {
	r := Runargs{}
	return &r
}
