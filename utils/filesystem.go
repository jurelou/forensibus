package utils

import (
	"os"
	// "fmt"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/h2non/filetype"
)

type FindFilesParams struct {
	Path string
	// Recurse			bool
	// Type		string
	PathPatterns []string
	FileFormats  []string
}

func FindFiles(params FindFilesParams) ([]string, error) {
	var pathRegexes []*regexp.Regexp
	var files []string
	var latestError error = nil

	checkFileFormat := (len(params.FileFormats) != 0)
	fileHead := make([]byte, 261)

	addPath := func(filePath string) error {
		// No conditions provided, add the current file

		if len(pathRegexes) == 0 && !checkFileFormat {
			files = append(files, filePath)
			return nil
		}

		// Check if file matches the given mime types
		if checkFileFormat {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()
			file.Read(fileHead)
			for _, mime := range params.FileFormats {
				if filetype.IsMIME(fileHead, mime) {
					files = append(files, filePath)
					return nil
				}
			}
		}

		// Check if file has a valid path
		for _, rex := range pathRegexes {
			if rex.MatchString(filePath) {
				files = append(files, filePath)
				return nil
			}
		}
		return nil
	}

	// Compile file paths patterns
	for _, pattern := range params.PathPatterns {
		compiledRegex, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		pathRegexes = append(pathRegexes, compiledRegex)

	}

	// Check if it's a folder
	fileInfo, errStat := os.Stat(params.Path)
	if errStat != nil {
		return nil, errStat
	}

	if fileInfo.IsDir() {
		filepath.Walk(params.Path, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			latestError = addPath(path)
			return nil
		})
	} else {
		latestError = addPath(params.Path)
	}
	return files, latestError
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func GetCurrentHomeDirectory() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}

func ConvertRelativePath(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		home, err := GetCurrentHomeDirectory()
		if err != nil {
			return path, err
		}
		return filepath.Join(home, path[1:]), nil
	} else {
		abs, err := filepath.Abs(path)
		if err != nil {
			return path, err
		}
		return abs, nil
	}
}

// func init() {
// 	fmt.Println("hello here")
// }
