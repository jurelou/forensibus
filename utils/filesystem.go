package utils

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

var buffSize = 1024 * 10

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

	checkFileFormat := (len(params.FileFormats) != 0)

	addPath := func(filePath string) {
		// No conditions provided, add the current file

		if len(pathRegexes) == 0 && !checkFileFormat {
			files = append(files, filePath)
			return
		}
		magicMatches := false
		// Check if file matches the given mime types
		if checkFileFormat {
			magic := GetMagic(filePath)
			for _, mime := range params.FileFormats {
				if strings.Contains(magic, mime) {
					magicMatches = true
					break
				}
			}
		}

		// Check if file has a valid path
		for _, rex := range pathRegexes {
			if rex.MatchString(filePath) {
				if !checkFileFormat || (checkFileFormat && magicMatches) {
					files = append(files, filePath)
					return
				}
				return
			}
		}
	}

	// Compile file paths patterns
	for _, pattern := range params.PathPatterns {
		compiledRegex, err := regexp.Compile(pattern)
		if err != nil {
			return nil, fmt.Errorf("could not compile pattern %s: %w", pattern, err)
		}
		pathRegexes = append(pathRegexes, compiledRegex)

	}

	// Check if it's a folder
	fileInfo, errStat := os.Stat(params.Path)
	if errStat != nil {
		return nil, fmt.Errorf("could not stat %s: %w", params.Path, errStat)
	}

	if fileInfo.IsDir() {
		err := filepath.Walk(params.Path, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			addPath(path)
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("could not walk path %s: %w", params.Path, err)
		}
	} else {
		addPath(params.Path)
	}
	return files, nil
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
		return "", fmt.Errorf("could not get current user: %w", err)
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
			return path, fmt.Errorf("could not get absolute path for %s: %w", path, err)
		}
		return abs, nil
	}
}

func CopyFile(in, out string) error {
	var destination *os.File
	source, err := os.Open(in)
	if err != nil {
		return fmt.Errorf("could not open %s: %w", in, err)
	}
	defer func() {
		if errClose := source.Close(); err != nil {
			err = errClose
		}
	}()

	outputInfo, err := os.Stat(out)
	if err != nil {
		// Output does not exists
		if err := os.MkdirAll(filepath.Dir(out), os.ModePerm); err != nil {
			return fmt.Errorf("could not create folder %s: %w", out, err)
		}
		destination, err = os.Create(out)
		if err != nil {
			return fmt.Errorf("could not create %s: %w", out, err)
		}
		defer func() {
			if errClose := destination.Close(); err != nil {
				err = errClose
			}
		}()
	} else {
		// Output already exists
		if outputInfo.IsDir() {
			out = filepath.Join(out, filepath.Base(in))
			destination, err = os.Create(out)
			if err != nil {
				return fmt.Errorf("could not create %s: %w", out, err)
			}
			defer func() {
				if errClose := destination.Close(); err != nil {
					err = errClose
				}
			}()

		} else {
			return fmt.Errorf("refusing to copy file %s to already existing file %s", in, out)
		}
	}
	buf := make([]byte, buffSize)
	for {
		numRead, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return fmt.Errorf("could not read from buffer: %w", err)
		}
		if numRead == 0 {
			break
		}
		if _, err := destination.Write(buf[:numRead]); err != nil {
			return fmt.Errorf("could not write to %s: %w", destination.Name(), err)
		}
	}
	return nil
}

// Create a directory from a given path recursively
// Skip if directory already exists
// func CreateDirectory(path string) error {
// 	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
// 		err := os.Mkdir(path, os.ModePerm)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
