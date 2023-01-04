package core

import (
	"fmt"

	// _ "github.com/jurelou/forensibus/processors/linux"
	// _ "github.com/jurelou/forensibus/processors/linux/commands"

	_ "github.com/jurelou/forensibus/processors/windows"
	_ "github.com/jurelou/forensibus/processors/windows/artifacts"
	_ "github.com/jurelou/forensibus/processors/windows/commands"
	"github.com/jurelou/forensibus/utils/processors"
)

func TotoAa() {
	fmt.Println("!!!")
	e, err := processors.Get("evtx")
	fmt.Println(e.Greet("salut"), err)
}
