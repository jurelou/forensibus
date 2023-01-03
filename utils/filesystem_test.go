package utils

import (
	_"fmt"
	"strings"
	"testing"
)

func TestFindFilesUnknownFolder(t *testing.T) {
	_, err := FindFiles(FindFilesParams{Path: "/this/should/not/exists"})
	if err == nil {
		t.Errorf("FindFiles should fail to find files from '/this/should/not/exists'")
	}
}

func TestFindFilesSingle(t *testing.T) {
	files, err := FindFiles(FindFilesParams{Path: "./filesystem.go"})

	if err != nil {
		t.Errorf("FindFiles returned an error: %s", err)
	}
	if len(files) != 1 || files[0] != "./filesystem.go" {
		t.Errorf("Invalid files found")
	}
}

func TestFindFilesSingleWithRegex(t *testing.T) {
	files, err := FindFiles(FindFilesParams{Path: "./filesystem.go", PathPatterns: []string{".*system.go"}})

	if err != nil {
		t.Errorf("FindFiles returned an error: %s", err)
	}
	if len(files) != 1 || files[0] != "./filesystem.go" {
		t.Errorf("Invalid files found")
	}
}

func TestFindFilesSingleWithInvalidRegex(t *testing.T) {
	files, err := FindFiles(FindFilesParams{Path: "./filesystem.go", PathPatterns: []string{".*nope.does.not.exists$"}})

	if err != nil {
		t.Errorf("FindFiles returned an error: %s", err)
	}
	if len(files) != 0 {
		t.Errorf("Invalid files found")
	}
}

func TestFindFilesFolder(t *testing.T) {
	files, err := FindFiles(FindFilesParams{Path: "."})

	if err != nil {
		t.Errorf("FindFiles returned an error: %s", err)
	}
	if len(files) < 2 {
		t.Errorf("Invalid files count found")
	}
}

func TestFindFilesFolderWithRegex(t *testing.T) {
	files, err := FindFiles(FindFilesParams{Path: ".", PathPatterns: []string{".*.this.should.not.exists", ".*filesystem(_test)?.go$"}})

	if err != nil {
		t.Errorf("FindFiles returned an error: %s", err)
	}
	if len(files) != 2 {
		t.Errorf("Invalid files count found")
	}
}

func TestFindFilesInvalidRegex(t *testing.T) {
	files, err := FindFiles(FindFilesParams{Path: ".", PathPatterns: []string{"invalid_file_name.nope"}})

	if err != nil {
		t.Errorf("FindFiles returned an error: %s", err)
	}
	if len(files) != 0 {
		t.Errorf("Invalid files count found")
	}
}


func TestFindFilesFromFolderWithFiletype(t *testing.T) {
	files, err := FindFiles(FindFilesParams{Path: "../", FileFormats: []string{"application/x-tar", "application/gzip"}})

	if err != nil {
		t.Errorf("FindFiles returned an error: %s", err)
	}
	if len(files) != 0 {
		t.Errorf("Invalid files count found")
	}
}

func TestConvertHomePath(t *testing.T) {
	in := "~/myfolder"
	out, err := ConvertRelativePath(in)
	if err != nil {
		t.Errorf("Error while converting relative path: %s", err)
	}
	if !strings.HasPrefix(out, "/home/") || !strings.HasSuffix(out, "/myfolder") {
		t.Errorf("Invalid path conversion")
	}
}

func TestConvertDotPath(t *testing.T) {
	in := "../utils/filesystem_test.go"
	out, err := ConvertRelativePath(in)
	if err != nil {
		t.Errorf("Error while converting relative path: %s", err)
	}
	if !strings.HasPrefix(out, "/") || !strings.HasSuffix(out, "/utils/filesystem_test.go") {
		t.Errorf("Invalid path conversion")
	}

	if !FileExists(out) {
		t.Errorf("filesystem_test.go should exists !")
	}
}

func TestConvertDotDotPath(t *testing.T) {
	in := "../../../../../../../../../../../../../../../../../../../../../tmp"
	out, err := ConvertRelativePath(in)
	if err != nil {
		t.Errorf("Error while converting relative path: %s", err)
	}
	if out != "/tmp" {
		t.Errorf("Invalid result for dot dot dot conversion")
	}
}

func TestConvertLocalPath(t *testing.T) {
	in := "filesystem_test.go"
	out, err := ConvertRelativePath(in)
	if err != nil {
		t.Errorf("Error while converting relative path: %s", err)
	}
	if !strings.HasPrefix(out, "/") || !strings.HasSuffix(out, "/utils/filesystem_test.go") {
		t.Errorf("Invalid path conversion")
	}

	if !FileExists(out) {
		t.Errorf("filesystem_test.go should exists !")
	}
}

func TestConvertLocalDotPath(t *testing.T) {
	in := "./filesystem_test.go"
	out, err := ConvertRelativePath(in)
	if err != nil {
		t.Errorf("Error while converting relative path: %s", err)
	}
	if !strings.HasPrefix(out, "/") || !strings.HasSuffix(out, "/utils/filesystem_test.go") {
		t.Errorf("Invalid path conversion")
	}

	if !FileExists(out) {
		t.Errorf("filesystem_test.go should exists !")
	}
}

func TestConvertStrangeDotPath(t *testing.T) {
	in := "./../utils/filesystem_test.go"
	out, err := ConvertRelativePath(in)
	if err != nil {
		t.Errorf("Error while converting relative path: %s", err)
	}
	if !strings.HasPrefix(out, "/") || !strings.HasSuffix(out, "/utils/filesystem_test.go") {
		t.Errorf("Invalid path conversion")
	}

	if !FileExists(out) {
		t.Errorf("filesystem_test.go should exists !")
	}
}