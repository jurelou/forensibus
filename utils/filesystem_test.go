package utils

import (
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