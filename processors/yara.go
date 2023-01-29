package common_processors

import (
	"fmt"

	yara "github.com/Velocidex/go-yara"

	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

type YaraProcessor struct{
	Compiler *yara.Compiler
}

func (proc YaraProcessor) Configure() error {
	compiler, err := yara.NewCompiler()
	if err != nil {
		return fmt.Errorf("Could not instanciate a yara compiler: %w", err)
	}
	proc.Compiler = compiler
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
