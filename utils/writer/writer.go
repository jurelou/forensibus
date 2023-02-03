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

func (w DefaultOutputWriter) Close() {
	return
}

func (w DefaultOutputWriter) SetTag(string) {
	return
}

func (w DefaultOutputWriter) SetDefaultIndex(string) {
	return
}

func (w DefaultOutputWriter) SetDefaultSource(string) {
	return
}

func (w DefaultOutputWriter) SetDefaultSourceType(string) {
	return
}

func (w DefaultOutputWriter) SetDefaultHost(string) {
	return
}

func New() OutputWriter {
	return NewHEC(utils.Config.Splunk.Hec.Address, utils.Config.Splunk.Hec.Token)
}
