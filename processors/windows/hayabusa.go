package windows_processors

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/google/uuid"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

type HayabusaProcessor struct {
	BinaryPath string
}

func (proc *HayabusaProcessor) Configure() error {
	hayabusaPath := "./external/hayabusa/"
	arch := runtime.GOARCH
	operatingSystem := runtime.GOOS
	switch os := operatingSystem; os {
	case "darwin":
		if arch == "arm" || arch == "arm64" {
			proc.BinaryPath = hayabusaPath + "hayabusa-2.1.0-mac-arm"
		} else if arch == "" {
			proc.BinaryPath = hayabusaPath + "hayabusa-2.1.0-mac-intel"
		}
	case "linux":
		proc.BinaryPath = hayabusaPath + "hayabusa-2.1.0-lin-musl"
	case "windows":
		if arch == "386" {
			proc.BinaryPath = hayabusaPath + "hayabusa-2.1.0-win-x86.exe"
		} else if arch == "amd64" {
			proc.BinaryPath = hayabusaPath + "hayabusa-2.1.0-win-x64.exe"
		}
	default:
		return fmt.Errorf("cannot run hayabusa on %s / %s platform", operatingSystem, arch)
	}
	if proc.BinaryPath == "" {
		return fmt.Errorf("could not find a supported hayabusa architecture for %s / %s", operatingSystem, arch)
	}

	version, err := getHayabusaVersion(proc.BinaryPath)
	if err != nil {
		return err
	}
	utils.Log.Infof("Configured hayabusa: %s", version)
	return nil
}

func (proc *HayabusaProcessor) Run(in string, config *processors.Config, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}

	evtxExt, exists := config.GetString("evtx_extension")
	if !exists {
		errors.Add(fmt.Errorf("missing hayabusa config: evtx_extension"))
		return errors
	}
	outFile, err := proc.RunHayabusa(in, evtxExt, out.GetTag())
	if err != nil {
		errors.Add(err)
		return errors
	}
	if _, err := os.Stat(outFile); err != nil {
		errors.Add(fmt.Errorf("hayabusa did not generate any files, probably because due to a lack of evtx files"))
		return errors
	}

	fd, err := os.Open(outFile)
	if err != nil {
		errors.Add(err)
		return errors
	}

	fileScanner := bufio.NewScanner(fd)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		e := writer.NewEvent(fileScanner.Text())
		out.WriteEvent(e)
	}

	fd.Close()

	return errors
}

func (proc *HayabusaProcessor) RunHayabusa(in, evtxExtension, tag string) (string, error) {
	isDir, err := IsDirectory(in)
	if err != nil {
		return "", fmt.Errorf("invalid input %s: %w", in, err)
	}
	outputFile := filepath.Join(utils.Config.TempFolder, fmt.Sprintf("hayabusa_%s_%s.json", tag, uuid.NewString()))
	args := []string{
		"json-timeline",
		"--quiet-errors",
		"--jsonl",
		"--ISO-8601",
		"--no-summary",
		"--profile",
		"super-verbose",
		"--rules-config",
		"./external/hayabusa/rules/config/",
		"--target-file-ext",
		evtxExtension,
		"--rules",
		utils.Config.SigmaRulesFolder,
		"--output",
		outputFile,
	}
	if isDir {
		args = append(args, "--directory", in)
	} else {
		args = append(args, "--file", in)
	}
	utils.Log.Debugf("Running hayabusa command: %v", args)
	out, err := execve(proc.BinaryPath, args)
	utils.Log.Debug(out.String())
	if err != nil {
		return "", fmt.Errorf("error while running hayabusa: %s", out.String())
	}
	return outputFile, nil
}

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

func execve(bin string, args []string) (bytes.Buffer, error) {
	var out bytes.Buffer
	cmd := exec.Command(bin, args...)
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		return out, err
	}
	return out, nil
}

func getHayabusaVersion(binaryPath string) (string, error) {
	if _, err := os.Stat(binaryPath); err != nil {
		return "", fmt.Errorf("could not find hayabusa binary %s", binaryPath)
	}
	out, err := execve(binaryPath, []string{})
	if err != nil {
		return "", fmt.Errorf("could not check hayabusa version: %w", err)
	}
	stdout := strings.Split(strings.ReplaceAll(out.String(), "\r\n", "\n"), "\n")
	if len(stdout) < 4 {
		return "", fmt.Errorf("invalid return from hayabusa, expected at least 4 lines: %s", out.String())
	}
	return stdout[2], nil
}

func init() {
	processors.Register("sigma", &HayabusaProcessor{})
}
