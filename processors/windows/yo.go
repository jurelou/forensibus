package windows

import (
	"fmt"
	_"github.com/jurelou/forensibus/utils/processors"
)

type Component1 struct {
	// Input string `yaml:"input"`
}

func (c Component1) Greet(s string) string {
	return fmt.Sprintf(
		"Hello, the input you provided to me was: %s", s)
}

// func (c Component1) UnmarshalSettings(settings []byte) processors.Processor {
// 	var c1 Component1
// 	yaml.Unmarshal(settings, &c1)
// 	return c1
// }

func init() {
	// processors.Register("salaaut", Component1{})
}
