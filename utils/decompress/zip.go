package decompress

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func DecompressZip(in string, out string) error {
	archive, err := zip.OpenReader(in)
	if err != nil {
		return err
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(out, f.Name)
		if !strings.HasPrefix(filePath, filepath.Clean(out)+string(os.PathSeparator)) {
			return fmt.Errorf("refuse to decompress %s (probably because of ..)", f.Name)
		}

		if f.FileInfo().IsDir() {
			if err := os.Mkdir(filePath, os.ModePerm); err != nil {
				return fmt.Errorf("could not create folder %s: %w", filePath, err)
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}

		dstFile.Close()
		fileInArchive.Close()
	}
	return nil
}
