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

type DefaultOutputWriter struct{}

func (DefaultOutputWriter) Close() {}

func (DefaultOutputWriter) SetTag(_ string) {}

func (DefaultOutputWriter) SetDefaultIndex(_ string) {}

func (DefaultOutputWriter) SetDefaultSource(_ string) {}

func (DefaultOutputWriter) SetDefaultSourceType(_ string) {}

func (DefaultOutputWriter) SetDefaultHost(_ string) {}

func New() OutputWriter {
	return NewHEC(utils.Config.Splunk.Hec.Address, utils.Config.Splunk.Hec.Token)
}
