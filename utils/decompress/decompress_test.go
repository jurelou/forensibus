package decompress_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/decompress"
)

func TestDecompressToFile(t *testing.T) {
	_, err := decompress.Decompress("../../dataset/archives/Simple.zip", "../../dataset/files/empty.txt", []string{})
	if err == nil {
		t.Errorf("Decompress to a file should fail")
	}
}

func TestDecompressUnknownFile(t *testing.T) {
	_, err := decompress.Decompress("/tmp/.this_does_does_exists", "/tmp/out", []string{})
	if err == nil {
		t.Errorf("Decompress should fail on invalid file")
	}
}

func TestDecompressInvalidFile(t *testing.T) {
	_, err := decompress.Decompress("../../datasets/pipelines/nested.hcl", "/tmp/out", []string{})
	if err == nil {
		t.Errorf("Decompress text file should fail")
	}
	fmt.Println()
}

func TestDecompressSameName(t *testing.T) {
	outputFolder := "/tmp/.forensibus_test_duplicate"
	var out string
	err := os.RemoveAll(outputFolder)
	if err != nil {
		t.Fatalf("Could not remove %s", outputFolder)
	}
	// 1st round
	out, err = decompress.Decompress("../../datasets/archives/Simple.zip", outputFolder, []string{})
	if err != nil {
		t.Errorf("Error while decompressing Simple.zip: %s", err.Error())
	}
	if out != outputFolder {
		t.Errorf("Invalid output folder: %s", out)
	}
	outputFile := filepath.Join(outputFolder, "somefolder/anotherfile")
	exists := utils.FileExists(outputFile)
	if !exists {
		t.Fatalf("file %s should exists", outputFile)
	}

	// 2nd round
	out, err = decompress.Decompress("../../datasets/archives/Simple.zip", outputFolder, []string{})
	if err != nil {
		t.Errorf("Error while decompressing Simple.zip a second time: %s", err.Error())
	}
	if out != filepath.Join(outputFolder, "Simple") {
		t.Errorf("Invalid output folder: %s", out)
	}
	outputFile = filepath.Join(outputFolder, "Simple/somefolder/anotherfile")
	exists = utils.FileExists(outputFile)
	if !exists {
		t.Fatalf("file %s should exists", outputFile)
	}

	// 3rd round
	out, err = decompress.Decompress("../../datasets/archives/Simple.zip", outputFolder, []string{})
	if err != nil {
		t.Errorf("Error while decompressing Simple.zip a third time: %s", err.Error())
	}
	if !strings.HasPrefix(out, filepath.Join(outputFolder, "Simple")) {
		t.Errorf("Invalid output folder: %s", out)
	}
	files, err := os.ReadDir(outputFolder)
	if err != nil {
		t.Fatalf("Error while reading folder %s", outputFolder)
	}
	if len(files) != 4 {
		t.Fatalf("3rd decompression did not succeed")
	}
}

func TestDecompress7z(t *testing.T) {
	outputFolder := "/tmp/.forensibus_decompress_sevenzip"
	outputFile := filepath.Join(outputFolder, "test7z/testfile")

	err := os.RemoveAll(outputFolder)
	if err != nil {
		t.Fatalf("Could not remove %s", outputFolder)
	}
	_, err = decompress.Decompress("../../datasets/archives/Simple.7z", outputFolder, []string{})
	if err != nil {
		t.Errorf("Error while decompressing Simple.7z: %s", err.Error())
	}
	if exists := utils.FileExists(outputFile); !exists {
		t.Fatalf("7z archive did not successfully decompressed files")
	}
	content, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Could not read %s", outputFile)
	}
	if string(content) != "it works!\n" {
		t.Fatalf("output file `testfile` content differs")
	}
}

func TestDecompressZip(t *testing.T) {
	outputFolder := "/tmp/.forensibus_decompress_zip"
	outputFile := filepath.Join(outputFolder, "somefolder/anotherfile")

	err := os.RemoveAll(outputFolder)
	if err != nil {
		t.Fatalf("Could not remove %s", outputFolder)
	}
	_, err = decompress.Decompress("../../datasets/archives/Simple.zip", outputFolder, []string{})
	if err != nil {
		t.Errorf("Error while decompressing Simple.7z: %s", err.Error())
	}

	if exists := utils.FileExists(outputFile); !exists {
		t.Fatalf("zip archive did not successfully decompressed files")
	}
	content, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Could not read %s", outputFile)
	}
	if string(content) != "anotherfile content\n" {
		t.Fatalf("output file `testfile` content differs")
	}
}

func TestDecompressZipPasswd(t *testing.T) {
	outputFolder := "/tmp/.forensibus_decompress_zip_passwd"
	outputFile := filepath.Join(outputFolder, "somefolder/anotherfile")

	err := os.RemoveAll(outputFolder)
	if err != nil {
		t.Fatalf("Could not remove %s", outputFolder)
	}

	_, err = decompress.Decompress("../../datasets/archives/SimplePasswd.zip", outputFolder, []string{"aa", "passwd", "bb"})
	if err != nil {
		t.Errorf("Error while decompressing SimplePasswd.zip: %s", err.Error())
	}

	if exists := utils.FileExists(outputFile); !exists {
		t.Fatalf("zip archive did not successfully decompressed files")
	}
	content, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Could not read %s", outputFile)
	}
	if string(content) != "anotherfile content\n" {
		t.Fatalf("output file `testfile` content differs")
	}
}

func TestDecompressInvalidZipPasswd(t *testing.T) {
	outputFolder := "/tmp/.forensibus_decompress_zip_passwd"

	err := os.RemoveAll(outputFolder)
	if err != nil {
		t.Fatalf("Could not remove %s", outputFolder)
	}

	_, err = decompress.Decompress("../../datasets/archives/SimplePasswd.zip", outputFolder, []string{"invalid", "passwordgiven"})
	if err == nil {
		t.Errorf("Should fail to decompress with invalid passwd")
	}

	if string(err.Error()) != "could not decrypt archive ../../datasets/archives/SimplePasswd.zip with passwords [invalid passwordgiven]" {
		t.Fatalf("Invalid error message: %s", err.Error())
	}
}
