package common_processors

import (
	"encoding/json"

	"gopkg.in/ini.v1"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

type IniProcessor struct {
	processors.Default
}

func (IniProcessor) Run(in string, _ *processors.Config, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}
	utils.Log.Infof("Running ini processor against %s", in)

	cfg, err := ini.Load(in)
	if err != nil {
		errors.Add(err)
		return errors
	}
	content := map[string]interface{}{}

	for _, sexName := range cfg.SectionStrings() {
		sex, _ := cfg.GetSection(sexName)
		if sexName == "DEFAULT" {
			for _, value := range sex.Keys() {
				content[value.Name()] = value.Value()
			}
		} else {
			content[sexName] = map[string]string{}
			for _, value := range sex.Keys() {
				content[sexName].(map[string]string)[value.Name()] = value.Value()
			}
		}

	}
	serialized, err := json.Marshal(content)
	if err != nil {
		errors.Add(err)
		return errors
	}

	e := writer.NewEvent(string(serialized))
	out.WriteEvent(e)
	return errors
}

func init() {
	processors.Register("ini", &IniProcessor{})
}
