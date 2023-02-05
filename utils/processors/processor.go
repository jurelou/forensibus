package processors

import (
	"github.com/jurelou/forensibus/utils/writer"
)

type Processor interface {
	Configure() error
	Run(string, *Config, writer.OutputWriter) PError
}

type Default struct{}

func (Default) Configure() error {
	return nil
}

type Config struct {
	RawConfig map[string]string
}

func (conf Config) GetString(key string) (string, bool) {
	val, exists := conf.RawConfig[key]
	if exists {
		return val, true
	}
	return "", false
}
