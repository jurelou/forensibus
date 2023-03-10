package common_processors

import (
	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

type DummyProcessor struct {
	// processors.Default
}

func (DummyProcessor) Configure() error {
	return nil
}

func (DummyProcessor) Run(in string, _ *processors.Config, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}
	utils.Log.Infof("Running dummy processor against %s", in)
	return errors
}

func init() {
	processors.Register("dummy", &DummyProcessor{})
}
