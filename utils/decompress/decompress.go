package decompress

import (
	"errors"
	"fmt"
	"os"
	_ "path"
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
			return "", err
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
		return "", errors.New(fmt.Sprintf("Cannot extract %s to file %s", out, in))

	} else {
		// Folder does not exists, create it
		if err := os.MkdirAll(out, os.ModePerm); err != nil {
			return "", err
		}
	}
	return out, nil
}

func Decompress(in string, out string) (string, error) {
	if exists := utils.FileExists(in); !exists {
		return "", errors.New(fmt.Sprintf("File %s does not exists.", in))
	}

	outputFolder, err := genOutputFolder(in, out)
	if err != nil {
		return "", err
	}

	magic := utils.GetMagic(in)

	if magic == "application/x-7z-compressed" {
		return outputFolder, DecompressSevenZip(in, outputFolder)
		// } else if filetype.IsMIME(head, "application/x-tar") {
		// 	return DecompressSevenZip(in, out)
		// } else if filetype.IsMIME(head, "application/x-xz") {
		// 	return DecompressSevenZip(in, out)
		// } else if filetype.IsMIME(head, "application/vnd.rar") {
		// 	return DecompressSevenZip(in, out)
		// } else if filetype.IsMIME(head, "application/x-bzip2") {
		// 	return DecompressSevenZip(in, out)
		// } else if filetype.IsMIME(head, "application/x-tar") {
		// 	return DecompressSevenZip(in, out)
	} else if magic == "application/zip" {
		return outputFolder, DecompressSevenZip(in, outputFolder)
	}

	// handle := magic.NewMagicHandle(magic.MAGIC_NONE)

	return outputFolder, errors.New(fmt.Sprintf("Invalid archive type"))
}
