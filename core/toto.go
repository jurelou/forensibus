package core

import (
	"fmt"
)


func Yo(pipelineconfigFile string, paths []string) {
	fmt.Println("hello world", paths, "===", pipelineconfigFile)
	config, err := LoadDSLFile(pipelineconfigFile); if err != nil {
		fmt.Println(err,config)
	}
}