package decompress

import (
	// "errors"
	"archive/zip"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	// "github.com/jurelou/forensibus/utils"
)

func DecompressZip(in string, out string) error {
	archive, err := zip.OpenReader(in)
	if err != nil {
		return err
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(out, f.Name)
		fmt.Println(f.Name)
		if !strings.HasPrefix(filePath, filepath.Clean(out)+string(os.PathSeparator)) {
			return errors.New(fmt.Sprintf("Invalid zip file %s", f.Name))
		}

		if f.FileInfo().IsDir() {
			os.Mkdir(filePath, os.ModePerm)
			continue
		}
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

	}
	return nil
}
