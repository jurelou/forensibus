package common_processors

import (
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"

	yara "github.com/Velocidex/go-yara"
)

type YaraProcessor struct{}

func (proc YaraProcessor) Configure() error {
	return nil
}

func (proc YaraProcessor) Run(in string, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}

	yara.NewCompiler()
	return errors
}

func init() {
	processors.Register("yara", YaraProcessor{})
}
