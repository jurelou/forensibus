package utils

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	buffSize = 1024 * 10
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

	addPath := func(filePath string) error {
		// No conditions provided, add the current file

		if len(pathRegexes) == 0 && !checkFileFormat {
			files = append(files, filePath)
			return nil
		}

		// Check if file matches the given mime types
		if checkFileFormat {
			magic := GetMagic(filePath)

			for _, mime := range params.FileFormats {
				if mime == magic {
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

func CopyFile(in, out string) error {
	var destination *os.File
	source, err := os.Open(in)
	if err != nil {
		return err
	}
	defer source.Close()

	outputInfo, err := os.Stat(out)
	if err != nil {
		// Output does not exists
		if err := os.MkdirAll(filepath.Dir(out), os.ModePerm); err != nil {
			return err
		}
		destination, err = os.Create(out)
		if err != nil {
			return err
		}
		defer destination.Close()
	} else {
		// Output already exists
		if outputInfo.IsDir() {
			out = filepath.Join(out, filepath.Base(in))
			destination, err = os.Create(out)
			if err != nil {
				return err
			}
			defer destination.Close()
		} else {
			return errors.New(fmt.Sprintf("Refusing to copy file %s to already existing file %s", in, out))
		}
	}
	buf := make([]byte, buffSize)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}
