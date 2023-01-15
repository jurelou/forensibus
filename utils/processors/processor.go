package processors

import "github.com/jurelou/forensibus/utils/writer"
type Processor interface {
	Configure() error
	Run(string, writer.OutputWriter) error
	// UnmarshalSettings([]byte) Processor
}
