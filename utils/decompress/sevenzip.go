package decompress

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/jurelou/forensibus/utils"
)

var (
	sevenZipPath = ""
)

type Entry struct {
	Path string
	Size int64
	// PackedSize int
	// Modified   time.Time
	// Attributes string
	// CRC        string
	Encrypted string
	Method    string
	// Block      int
}

func findSevenZipFromPath() string {
	if p, err := exec.LookPath("7z"); err == nil {
		return p
	}
	return ""
}

func findSevenZipFromCurrentFolder() string {
	var path string
	exePath, err := os.Executable()
	if err != nil {
		return path
	}
	exeDirname := filepath.Dir(exePath)

	switch runtime.GOOS {
	case "windows":
		path = filepath.Join(exeDirname, "external/7z/windows/7z.exe")
	case "linux":
		path = filepath.Join(exeDirname, "external/7z/linux/7zz")
	case "darwin":
		path = filepath.Join(exeDirname, "external/7z/darwin/7zz")
	}
	if _, err := os.Stat(path); err != nil {
		return ""
	}
	return path
}

func getSevenZipPath() {
	sevenZipPath = findSevenZipFromPath()
	if sevenZipPath != "" {
		return
	}
	sevenZipPath = findSevenZipFromCurrentFolder()
}

func advanceToFirstEntry(scanner *bufio.Scanner) error {
	for scanner.Scan() {
		s := scanner.Text()
		if s == "----------" {
			return nil
		}
	}
	err := scanner.Err()
	if err == nil {
		return errors.New("Could not find any entries")
	}
	return err
}

func getEntryLines(scanner *bufio.Scanner) ([]string, error) {
	var res []string
	for scanner.Scan() {
		s := scanner.Text()
		s = strings.TrimSpace(s)
		if s == "" {
			break
		}
		res = append(res, s)
	}
	err := scanner.Err()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func checkLines(lines []string) (bool, error) {

	for _, s := range lines {
		parts := strings.SplitN(s, " =", 2)
		if len(parts) != 2 {
			return false, fmt.Errorf("unexpected line: '%s'", s)
		}
		name := strings.ToLower(parts[0])
		value := strings.TrimSpace(parts[1])
		if value == "" {
			continue
		}
		if name == "encrypted" && value != "-" {
			return true, nil
		}

	}
	return false, nil
}

func findEncryptedFlag(data []byte) (bool, error) {
	buff := bytes.NewBuffer(data)
	scanner := bufio.NewScanner(buff)
	err := advanceToFirstEntry(scanner)
	if err != nil {
		return false, err
	}

	for {
		lines, err := getEntryLines(scanner)
		if err != nil {
			return false, err
		}
		if len(lines) == 0 {
			break
		}
		isEncrypted, err := checkLines(lines)
		if err != nil {
			return false, err
		}
		if isEncrypted {
			return true, nil
		}
	}
	return false, nil
}

func IsEncrypted(in string) (bool, error) {
	params := []string{"l", "-slt", "-sccUTF-8", in}
	cmd := exec.Command(sevenZipPath, params...)
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}
	return findEncryptedFlag(out)
}

func decompress(in string, out string, password string) error {
	params := []string{"x", fmt.Sprintf("-o%s", out), in}
	if password != "" {
		params = append(params, fmt.Sprintf("-p%s", password))
	}
	cmd := exec.Command(sevenZipPath, params...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return errors.New(fmt.Sprintf("Error while decompressing 7z %s: %s", in, string(stderr.Bytes())))
	}
	return nil
}

func DecompressSevenZip(in string, out string) error {
	fileFinfo, errStat := os.Stat(in)
	if errStat != nil {
		return errStat
	}
	if fileFinfo.IsDir() {
		return errors.New(fmt.Sprintf("Cannot decompress folder %s ..", in))
	}

	getSevenZipPath()
	if sevenZipPath == "" {
		return errors.New("Could not find 7z")
	}

	encrypted, errEncrypted := IsEncrypted(in)
	if errEncrypted != nil {
		return errEncrypted
	}

	if encrypted == false {
		// Decompress unencrypted archive
		if err := decompress(in, out, ""); err != nil {
			fmt.Printf("Could not decompress %s: %s", in, err.Error())
			return err
		}
		return nil
	}

	// Archive is encrypted
	if len(utils.Config.ArchivePasswords) == 0 {
		return errors.New(fmt.Sprintf("Archive %s is encrypted, but no passwords are provided.", in))
	}
	for _, passwd := range utils.Config.ArchivePasswords {
		if err := decompress(in, out, passwd); err == nil {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Could not decrypt archive %s whith passwords %s", in, utils.Config.ArchivePasswords))
}
