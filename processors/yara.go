package common_processors

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hillu/go-yara/v4"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

type YaraProcessor struct {
	Compiler *yara.Compiler
}

func (proc *YaraProcessor) Configure() error {
	c, err := yara.NewCompiler()

	// c.DefineVariable("filepath", "")
	// c.DefineVariable("filename", "")
	// c.DefineVariable("extension", "")
	// c.DefineVariable("filetype", "")
	// c.DefineVariable("owner", "")

	if err != nil {
		return fmt.Errorf("could not initialize yara compiler: %w", err)
	}
	yaraFiles, err := utils.FindFiles(utils.FindFilesParams{
		Path:         "./external/yara_rules",
		PathPatterns: []string{".*.yara?$", ".*.yas$"},
	})
	if err != nil {
		return fmt.Errorf("could not list files in ./external/yara_rules: %w", err)
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
	// rules, err := c.GetRules()
	// if err != nil {
	// 	return fmt.Errorf("could not load rules: %w", err)
	// }
	proc.Compiler = c
	return nil
}

func (proc *YaraProcessor) Run(in string, _ *processors.Config, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}
	utils.Log.Infof("Running YARA processor against %s", in)

	compiler := proc.Compiler
	compiler.DefineVariable("filepath", in)
	compiler.DefineVariable("filename", filepath.Base(in))
	compiler.DefineVariable("extension", filepath.Ext(in))
	// compiler.DefineVariable("filetype", "")
	// compiler.DefineVariable("owner", "")

	rules, err := compiler.GetRules()
	if err != nil {
		errors.Add(fmt.Errorf("could not load rules: %w", err))
		return errors

	}

	scanner, err := yara.NewScanner(rules)
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
