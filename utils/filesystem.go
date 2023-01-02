package utils

import (
	"fmt"
	_ "errors"
	"os"
	"regexp"
)
type FindFilesParams struct {
	Path string
	PathPatterns []string
	FileMagics []string
}

func FindFiles(params FindFilesParams) ([]string, error) {
	fmt.Println("!!!", params)
	var pathRegexes []*regexp.Regexp
	var files []string

	// Open the given file or folder
	// filePtr, errOpen := os.Open(params.Path); if errOpen != nil {
	// 	return nil, errOpen
	// }
	// defer filePtr.Close()
	fileInfo, errStat := os.Stat(params.Path); if errStat != nil {
		return nil, errStat
	}
	if fileInfo.IsDir() {
		fmt.Println("its a folder")
	}

	// Compile file paths patterns
	for _, pattern := range params.PathPatterns {
		compiledRegex, err := regexp.Compile(pattern); if err != nil {
			return nil, err
		}
		pathRegexes = append(pathRegexes, compiledRegex)

	}
	fmt.Println(pathRegexes)
	return files, nil
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
