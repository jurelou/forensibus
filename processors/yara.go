package common_processors

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	yara "github.com/Velocidex/go-yara"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

type YaraProcessor struct {
	Rules *yara.Rules
}

func (proc *YaraProcessor) Configure() error {
	c, err := yara.NewCompiler()
	if err != nil {
		return fmt.Errorf("could not initialize yara compiler: %w", err)
	}
	yaraFiles, err := utils.FindFiles(utils.FindFilesParams{
		Path:         utils.Config.YaraRulesFolder,
		PathPatterns: []string{".*.yara?$", ".*.yas$"},
	})
	if err != nil {
		return fmt.Errorf("could not list files in %s: %w", utils.Config.YaraRulesFolder, err)
	}
	for _, file := range yaraFiles {
		f, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("could not open yara rule %s: %w", file, err)
		}
		namespace := filepath.Base(file)
		err = c.AddFile(f, namespace)
		if err != nil {
			return fmt.Errorf("could not load rule %s: %w", file, err)
		}
		f.Close()
	}
	rules, err := c.GetRules()
	if err != nil {
		return fmt.Errorf("could not load rules: %w", err)
	}
	proc.Rules = rules
	return nil
}

func (proc *YaraProcessor) Run(in string, _ *processors.Config, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}
	utils.Log.Infof("Running YARA processor against %s", in)

	scanner, err := yara.NewScanner(proc.Rules)
	if err != nil {
		errors.Add(err)
		return errors
	}

	var m yara.MatchRules
	err = scanner.SetCallback(&m).SetFlags(yara.ScanFlagsFastMode).ScanFile(in)
	if err != nil {
		errors.Add(err)
		return errors
	}

	// TODO: make a custom matcher: https://github.com/Velocidex/velociraptor/blob/4ebef17438a08eaa5b49b5ca8f9f006669b41d12/vql/common/yara.go#L417

	for _, match := range m {
		jsonMatch, err := json.Marshal(match)
		if err != nil {
			errors.Add(err)
			continue
		}
		e := writer.NewEvent(string(jsonMatch))
		out.WriteEvent(e)
	}
	return errors
}

func init() {
	processors.Register("yara", &YaraProcessor{})
}
