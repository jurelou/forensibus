package common_processors

import (
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)


type DummyProcessor struct {
}

func (proc DummyProcessor) Configure() error {
	return nil
}

func (proc DummyProcessor) Run(in string, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}

	return errors
}

func init() {
	processors.Register("dummy", DummyProcessor{})
}
