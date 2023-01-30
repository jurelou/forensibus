package decompress

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"

	"github.com/jurelou/forensibus/utils"
)

func createNamedFolder(filePath string) (string, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		// Folder does not exists, create it
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return "", fmt.Errorf("could not create folder %s: %w", filePath, err)
		}
		return filePath, nil
	} else {
		id := uuid.New()
		return createNamedFolder(fmt.Sprintf("%s_%s", filePath, id.String()[:4]))
	}
}

func genOutputFolder(in string, out string) (string, error) {
	fileInfo, err := os.Stat(out)
	if err == nil {
		// Folder already exists, creates another folder with a suffix
		inputFilename := filepath.Base(in)
		inputFilenameWhithoutSuffix := strings.TrimSuffix(inputFilename, filepath.Ext(inputFilename))
		if fileInfo.IsDir() {
			return createNamedFolder(filepath.Join(out, filepath.Base(inputFilenameWhithoutSuffix)))
		}
		return "", fmt.Errorf("cannot extract %s to file %s", out, in)

	} else {
		// Folder does not exists, create it
		if err := os.MkdirAll(out, os.ModePerm); err != nil {
			return "", fmt.Errorf("could not create folder %s: %w", out, err)
		}
	}
	return out, nil
}

func Decompress(in string, out string) (string, error) {
	if exists := utils.FileExists(in); !exists {
		return "", fmt.Errorf("file %s does not exists.", in)
	}

	outputFolder, err := genOutputFolder(in, out)
	if err != nil {
		return "", err
	}

	magic := utils.GetMagic(in)

	if strings.Contains(magic, "7-zip archive") || strings.Contains(magic, "Zip archive") {
		return outputFolder, DecompressSevenZip(in, outputFolder)
	}
	// else if strings.Contains(magic, "Zip archive") {
	// 	return outputFolder, DecompressSevenZip(in, outputFolder)
	// }

	return outputFolder, fmt.Errorf("invalid archive type")
}
