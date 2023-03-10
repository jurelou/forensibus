package decompress_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/decompress"
)

func Test7zInvalidFile(t *testing.T) {
	err := decompress.DecompressSevenZip("./this_does_not_exists.xz", "/tmp/out", []string{})
	if err == nil {
		t.Fatalf("Decompress should fail on invalid file: `./this_does_not_exists.xz`")
	}
}

func Test7zSimple7z(t *testing.T) {
	file := "../../datasets/archives/Collect_empty.7z"
	outputFolder := "/tmp/.forensibus_test_7z"
	outputFile := filepath.Join(outputFolder, "empty_file")

	err := os.RemoveAll(outputFolder)
	if err != nil {
		t.Fatalf("Could not remove %s", outputFolder)
	}

	err = decompress.DecompressSevenZip(file, outputFolder, []string{})
	if err != nil {
		t.Fatalf("Decompressing an archive should not generate an error: %s", err.Error())
	}
	if exists := utils.FileExists(outputFile); !exists {
		t.Fatalf("7z archive did not successfully decompressed files")
	}
	content, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Could not read %s", outputFile)
	}
	if string(content) != "this file is empty\n" {
		t.Fatalf("output file `empty_file` content differs")
	}
}

func TestEncrypted7z(t *testing.T) {
	file := "../../datasets/archives/EmptyEncryptedPasswdIsPasswd.7z"
	outputFolder := "/tmp/.forensibus_test_7z_passwd"
	outputFile := filepath.Join(outputFolder, "empty.exe")

	err := os.RemoveAll(outputFolder)
	if err != nil {
		t.Fatalf("Could not remove %s", outputFolder)
	}

	err = decompress.DecompressSevenZip(file, outputFolder, []string{"good", "passwd"})
	if err != nil {
		t.Fatalf("Decompressing an encrypted archive should not generate an error")
	}
	if exists := utils.FileExists(outputFile); !exists {
		t.Fatalf("Encrypted archive did not successfully decompressed files")
	}
	content, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Could not read %s", outputFile)
	}
	if string(content) != "its an empty file\n" {
		t.Fatalf("output file `empty.exe` content differs")
	}
}

func TestEncrypted7zInvalidPasswd(t *testing.T) {
	file := "../../datasets/archives/EmptyEncryptedPasswdIsPasswd.7z"
	outputFolder := "/tmp/.forensibus_test_7z_invalid_passwd"
	outputFile := filepath.Join(outputFolder, "empty.exe")

	err := os.RemoveAll(outputFolder)
	if err != nil {
		t.Fatalf("Could not remove %s", outputFolder)
	}

	err = decompress.DecompressSevenZip(file, outputFolder, []string{"invalid", "passw00rd"})
	if err == nil {
		t.Fatalf("Decompressing an encrypted archive with invalid passwords should fail")
	}
	if utils.FileExists(outputFile) {
		t.Fatalf("On failure, should not generate output files")
	}
}

func TestEncrypted7zWhithoutPasswords(t *testing.T) {
	file := "../../datasets/archives/EmptyEncryptedPasswdIsPasswd.7z"
	err := decompress.DecompressSevenZip(file, "/tmp/nope", []string{})
	if err == nil {
		t.Fatalf("Decompressing an encrypted archive whithout passwords should generate an error")
	}
	if !strings.Contains(err.Error(), "is encrypted, but no passwords were provided") {
		t.Fatalf("Invalid error while decompressing encrypted archive")
	}
}
