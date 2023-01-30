package common_processors

import (
	yara "github.com/Velocidex/go-yara"

	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

type YaraProcessor struct {
	processors.Default
	Compiler *yara.Compiler
}

func (proc YaraProcessor) Run(in string, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}

	yara.NewCompiler()
	return errors
}

func init() {
	processors.Register("yara", YaraProcessor{})
}
