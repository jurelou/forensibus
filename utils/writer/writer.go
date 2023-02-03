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

func (DefaultOutputWriter) Close() {
	return
}

func (DefaultOutputWriter) SetTag(_ string) {
	return
}

func (DefaultOutputWriter) SetDefaultIndex(_ string) {
	return
}

func (DefaultOutputWriter) SetDefaultSource(_ string) {
	return
}

func (DefaultOutputWriter) SetDefaultSourceType(_ string) {
	return
}

func (DefaultOutputWriter) SetDefaultHost(_ string) {
	return
}

func New() OutputWriter {
	return NewHEC(utils.Config.Splunk.Hec.Address, utils.Config.Splunk.Hec.Token)
}
