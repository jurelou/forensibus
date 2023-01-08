package core

import (
	"fmt"
	"github.com/jurelou/forensibus/utils/processors"
)

func ListProcessors() {
	fmt.Println("Enabled processors:")

	for procName := range processors.Registry {
		fmt.Println("\t", procName)
	}
}
