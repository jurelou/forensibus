/*
Copyright © 2023 JURELOU
*/
package main

import (
	"os"

	"github.com/jurelou/forensibus/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
