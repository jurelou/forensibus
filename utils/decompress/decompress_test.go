package decompress

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/jurelou/forensibus/utils"
)

func TestDecompressUnknownFile(t *testing.T) {
	err := Decompress("/tmp/.this_does_does_exists", "/tmp/out")
	if err == nil {
		t.Errorf("Decompress should fail on invalid file")
	}
}

func TestDecompressInvalidFile(t *testing.T) {
	err := Decompress("../../datasets/pipelines/nested.hcl", "/tmp/out")
	if err == nil {
		t.Errorf("Decompress text file should fail")
	}
	fmt.Println()
}

func TestDecompress7z(t *testing.T) {
	outputFolder := "/tmp/.forensibus_decompress_sevenzip"
	outputFile := filepath.Join(outputFolder, "test7z/testfile")

	err := os.RemoveAll(outputFolder)
	if err != nil {
		t.Fatalf("Could not remove %s", outputFolder)
	}
	err = Decompress("../../datasets/archives/Simple.7z", outputFolder)
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
	outputFile := filepath.Join(outputFolder, "test7z/testfile")

	err := os.RemoveAll(outputFolder)
	if err != nil {
		t.Fatalf("Could not remove %s", outputFolder)
	}
	err = Decompress("../../datasets/archives/Simple.zip", outputFolder)
	if err != nil {
		t.Errorf("Error while decompressing Simple.7z: %s", err.Error())
	}

	fmt.Println(">>>>>>>", outputFile)
	// if exists := utils.FileExists(outputFile); !exists {
	// 	t.Fatalf("7z archive did not successfully decompressed files")
	// }
	// content, err := os.ReadFile(outputFile); if err != nil {
	// 	t.Fatalf("Could not read %s", outputFile)
	// }
	// if string(content) != "it works!\n" {
	// 	t.Fatalf("output file `testfile` content differs")
	// }
}
