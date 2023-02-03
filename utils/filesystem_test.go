package utils_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jurelou/forensibus/utils"
)

func TestFindFilesUnknownFolder(t *testing.T) {
	_, err := utils.FindFiles(utils.FindFilesParams{Path: "/this/should/not/exists"})
	if err == nil {
		t.Errorf("FindFiles should fail to find files from '/this/should/not/exists'")
	}
}

func TestFindFiles7Zip(t *testing.T) {
	files, err := utils.FindFiles(utils.FindFilesParams{Path: "../datasets", PathPatterns: []string{".*"}, FileFormats: []string{"application/x-7z-compressed"}})
	if err != nil {
		t.Errorf("FindFiles returned an error: %s", err)
	}
	for _, file := range files {
		if !strings.HasSuffix(file, ".7z") {
			t.Errorf(fmt.Sprintf("[PROBABLE FALSE NEGATIVE] Find with 7z mime type returned a file whithout .7z suffix: %s", file))
		}
	}
}

func TestFindFilesSingle(t *testing.T) {
	files, err := utils.FindFiles(utils.FindFilesParams{Path: "./filesystem.go"})
	if err != nil {
		t.Errorf("FindFiles returned an error: %s", err)
	}
	if len(files) != 1 || files[0] != "./filesystem.go" {
		t.Errorf("Invalid files found")
	}
}

func TestFindFilesSingleWithRegex(t *testing.T) {
	files, err := utils.FindFiles(utils.FindFilesParams{Path: "./filesystem.go", PathPatterns: []string{".*system.go"}})
	if err != nil {
		t.Errorf("FindFiles returned an error: %s", err)
	}
	if len(files) != 1 || files[0] != "./filesystem.go" {
		t.Errorf("Invalid files found")
	}
}

func TestFindFilesSingleWithInvalidRegex(t *testing.T) {
	files, err := utils.FindFiles(utils.FindFilesParams{Path: "./filesystem.go", PathPatterns: []string{".*nope.does.not.exists$"}})
	if err != nil {
		t.Errorf("FindFiles returned an error: %s", err)
	}
	if len(files) != 0 {
		t.Errorf("Invalid files found")
	}
}

func TestFindFilesFolder(t *testing.T) {
	files, err := utils.FindFiles(utils.FindFilesParams{Path: "."})
	if err != nil {
		t.Errorf("FindFiles returned an error: %s", err)
	}
	if len(files) < 2 {
		t.Errorf("Invalid files count found")
	}
}

func TestFindFilesFolderWithRegex(t *testing.T) {
	files, err := utils.FindFiles(utils.FindFilesParams{Path: ".", PathPatterns: []string{".*.this.should.not.exists", ".*filesystem(_test)?.go$"}})
	if err != nil {
		t.Errorf("FindFiles returned an error: %s", err)
	}
	if len(files) != 2 {
		t.Errorf("Invalid files count found")
	}
}

func TestFindFilesInvalidRegex(t *testing.T) {
	files, err := utils.FindFiles(utils.FindFilesParams{Path: ".", PathPatterns: []string{"invalid_file_name.nope"}})
	if err != nil {
		t.Errorf("FindFiles returned an error: %s", err)
	}
	if len(files) != 0 {
		t.Errorf("Invalid files count found")
	}
}

func TestFindFilesFromFolderWithFiletype(t *testing.T) {
	files, err := utils.FindFiles(utils.FindFilesParams{Path: "../core", FileFormats: []string{"application/x-tar", "application/gzip"}})
	if err != nil {
		t.Errorf("FindFiles returned an error: %s", err)
	}
	if len(files) != 0 {
		t.Errorf("Invalid files count found")
	}
}

func TestConvertHomePath(t *testing.T) {
	in := "~/myfolder"
	out, err := utils.ConvertRelativePath(in)
	if err != nil {
		t.Errorf("Error while converting relative path: %s", err)
	}
	if !strings.HasPrefix(out, "/home/") || !strings.HasSuffix(out, "/myfolder") {
		t.Errorf("Invalid path conversion")
	}
}

func TestConvertDotPath(t *testing.T) {
	in := "../utils/filesystem_test.go"
	out, err := utils.ConvertRelativePath(in)
	if err != nil {
		t.Errorf("Error while converting relative path: %s", err)
	}
	if !strings.HasPrefix(out, "/") || !strings.HasSuffix(out, "/utils/filesystem_test.go") {
		t.Errorf("Invalid path conversion")
	}

	if !utils.FileExists(out) {
		t.Errorf("filesystem_test.go should exists !")
	}
}

func TestConvertDotDotPath(t *testing.T) {
	in := "../../../../../../../../../../../../../../../../../../../../../tmp"
	out, err := utils.ConvertRelativePath(in)
	if err != nil {
		t.Errorf("Error while converting relative path: %s", err)
	}
	if out != "/tmp" {
		t.Errorf("Invalid result for ../../../ conversion")
	}
}

func TestConvertLocalPath(t *testing.T) {
	in := "filesystem_test.go"
	out, err := utils.ConvertRelativePath(in)
	if err != nil {
		t.Errorf("Error while converting relative path: %s", err)
	}
	if !strings.HasPrefix(out, "/") || !strings.HasSuffix(out, "/utils/filesystem_test.go") {
		t.Errorf("Invalid path conversion")
	}

	if !utils.FileExists(out) {
		t.Errorf("filesystem_test.go should exists !")
	}
}

func TestConvertLocalDotPath(t *testing.T) {
	in := "./filesystem_test.go"
	out, err := utils.ConvertRelativePath(in)
	if err != nil {
		t.Errorf("Error while converting relative path: %s", err)
	}
	if !strings.HasPrefix(out, "/") || !strings.HasSuffix(out, "/utils/filesystem_test.go") {
		t.Errorf("Invalid path conversion")
	}

	if !utils.FileExists(out) {
		t.Errorf("filesystem_test.go should exists !")
	}
}

func TestConvertStrangeDotPath(t *testing.T) {
	in := "./../utils/filesystem_test.go"
	out, err := utils.ConvertRelativePath(in)
	if err != nil {
		t.Errorf("Error while converting relative path: %s", err)
	}
	if !strings.HasPrefix(out, "/") || !strings.HasSuffix(out, "/utils/filesystem_test.go") {
		t.Errorf("Invalid path conversion")
	}

	if !utils.FileExists(out) {
		t.Errorf("filesystem_test.go should exists !")
	}
}

func TestCopyFileFile(t *testing.T) {
	output := "/tmp/.forensibus_test_copy_1"
	err := os.RemoveAll(output)
	if err != nil {
		t.Fatalf("Could not remove %s", output)
	}

	err = utils.CopyFile("../datasets/files/empty.txt", output)
	if err != nil {
		t.Errorf("Copy empty file should not fail: %s", err.Error())
	}

	if exists := utils.FileExists(output); !exists {
		t.Fatalf("Copy did not succeed to create output file")
	}
	content, err := os.ReadFile(output)
	if err != nil {
		t.Fatalf("Could not read %s: %s", output, err.Error())
	}
	if string(content) != "empty file\n" {
		t.Fatalf("output file content differs")
	}
	fmt.Println()
}

func TestCopyInvalidFile(t *testing.T) {
	output := "/tmp/.forensibus_test_copy_2"
	err := os.RemoveAll(output)
	if err != nil {
		t.Fatalf("Could not remove %s", output)
	}

	err = utils.CopyFile("/tmp/thi/does/not/exi", output)
	if err == nil {
		t.Errorf("Copy unknown file should fail")
	}
}

func TestCopyFileToFolder(t *testing.T) {
	output := "/tmp/.forensibus_test_copy_3"
	err := os.RemoveAll(output)
	if err != nil {
		t.Fatalf("Could not remove %s", output)
	}
	err = os.Mkdir(output, 0o700)
	if err != nil {
		t.Fatalf("Could not create %s", output)
	}
	err = utils.CopyFile("../datasets/files/empty.txt", output)
	if err != nil {
		t.Errorf("Copy empty file should not fail: %s", err.Error())
	}

	if exists := utils.FileExists(output); !exists {
		t.Fatalf("Copy did not succeed to create output file")
	}
	content, err := os.ReadFile(filepath.Join(output, "empty.txt"))
	if err != nil {
		t.Fatalf("Could not read %s: %s", output, err.Error())
	}
	if string(content) != "empty file\n" {
		t.Fatalf("output file content differs")
	}
	fmt.Println()
}

func TestCopyFileSubfolders(t *testing.T) {
	output := "/tmp/.forensibus_test_copy_4/this/is/a/nested/folder/myfile.txt"
	err := os.RemoveAll(output)
	if err != nil {
		t.Fatalf("Could not remove %s", output)
	}

	err = utils.CopyFile("../datasets/files/empty.txt", output)
	if err != nil {
		t.Errorf("Copy empty file should not fail: %s", err.Error())
	}

	if exists := utils.FileExists(output); !exists {
		t.Fatalf("Copy did not succeed to create output file")
	}
	content, err := os.ReadFile(output)
	if err != nil {
		t.Fatalf("Could not read %s: %s", output, err.Error())
	}
	if string(content) != "empty file\n" {
		t.Fatalf("output file content differs")
	}
	fmt.Println()
}
