package processors

import (
	"fmt"

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

func (conf Config) GetString(key string) (string, error) {
	val, exists := conf.RawConfig[key]
	if exists {
		return val, nil
	}
	return "", fmt.Errorf("config key %s does not exists", key)
}
