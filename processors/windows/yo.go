package windows

import (
	"fmt"
	"github.com/jurelou/forensibus/utils/processors"
	
)

type Component1 struct {
	Input string `yaml:"input"`
}

func (c Component1) Greet() string {
	return fmt.Sprintf(
		"Hello, the input you provided to me was: %s", c.Input)
}

// func (c Component1) UnmarshalSettings(settings []byte) processors.Processor {
// 	var c1 Component1
// 	yaml.Unmarshal(settings, &c1)
// 	return c1
// }

func init() {
	fmt.Println("hello winroot")
	processors.Register("salut", Component1{})
}