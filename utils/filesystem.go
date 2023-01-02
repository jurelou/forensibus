package utils

import (
	_ "fmt"
	_ "errors"
	"os"
	"regexp"
	"path/filepath"
)
type FindFilesParams struct {
	Path string
	PathPatterns []string
	FileMagics []string
}

func FindFiles(params FindFilesParams) ([]string, error) {
	var pathRegexes []*regexp.Regexp
	var files []string

	addFile := func (filePath string) {
		// No patterns provided, return
		if len(pathRegexes) == 0 {
			files = append(files, filePath)
			return
		}
		for _, rex := range pathRegexes {
			if rex.MatchString(filePath) {
				files = append(files, filePath)
				return
			}
		}
	}

	// Compile file paths patterns
	for _, pattern := range params.PathPatterns {
		compiledRegex, err := regexp.Compile(pattern); if err != nil {
			return nil, err
		}
		pathRegexes = append(pathRegexes, compiledRegex)

	}

	// Check if it's a folder
	fileInfo, errStat := os.Stat(params.Path); if errStat != nil {
		return nil, errStat
	}

	if fileInfo.IsDir() {
		filepath.Walk(params.Path, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			addFile(path)
			return nil
		})
	} else {
		addFile(params.Path)
	}

	return files, nil
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
