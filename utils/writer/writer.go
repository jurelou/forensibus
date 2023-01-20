package writer

import (
	"github.com/jurelou/forensibus/utils"
)

var Registry = make(map[string]OutputWriter)

type OutputWriter interface {
	// New()
	Close()
	SetTag(string)
	SetDefaultIndex(string)
	SetDefaultSource(string)
	SetDefaultSourceType(string)
	SetDefaultHost(string)

	WriteEvent(*Event)
}

func New() OutputWriter {
	return NewHEC(utils.Config.Splunk.Hec.Address, utils.Config.Splunk.Hec.Token)
}
