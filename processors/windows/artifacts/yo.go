package windows_artifacts

import (
	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
)

type EvtxProcessor struct {
	Input string `yaml:"input"`
}

func (p EvtxProcessor) Configure() error {
	return nil
}
func (p EvtxProcessor) Run(in string) error {
	utils.Log.Debugf("Run evtx processor against `%s`", in)
	return nil
}

func init() {
	processors.Register("evtx", EvtxProcessor{})
}
