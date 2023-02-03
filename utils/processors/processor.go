package processors

import "github.com/jurelou/forensibus/utils/writer"

type Processor interface {
	Configure() error
	Run(string, writer.OutputWriter) PError
}

type Default struct{}

func (Default) Configure() error {
	return nil
}
