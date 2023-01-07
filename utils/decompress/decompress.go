package decompress

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	_ "path"
	_ "path/filepath"
)

var (
	buffSize = 1024 * 10
)

func copyFile(in io.ReadCloser, out os.File) error {
	buf := make([]byte, buffSize)
	for {
		n, err := in.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("aa")
			return err
		}
		if n == 0 {
			break
		}
		if _, err := out.Write(buf[:n]); err != nil {
			fmt.Println("bb")
			return err
		}
	}
	return nil
}
func DecompressZip(filePath string) (string, error) {
	archive, err := zip.OpenReader(filePath)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer archive.Close()
	for _, f := range archive.File {
		fmt.Println(f.Name)
	}
	return "", nil
}
