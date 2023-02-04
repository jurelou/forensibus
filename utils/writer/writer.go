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

	GetTag() string
	WriteEvent(*Event)
}

type DefaultOutputWriter struct {
	Tag               string
	DefaultHost       string
	DefaultIndex      string
	DefaultSource     string
	DefaultSourceType string
}

func (*DefaultOutputWriter) Close() {}

func (out *DefaultOutputWriter) SetTag(tag string) {
	out.Tag = tag
}

func (out *DefaultOutputWriter) SetDefaultIndex(index string) {
	out.DefaultIndex = index
}

func (out *DefaultOutputWriter) SetDefaultSource(source string) {
	out.DefaultSource = source
}

func (out *DefaultOutputWriter) SetDefaultSourceType(sourceType string) {
	out.DefaultSourceType = sourceType
}

func (out *DefaultOutputWriter) SetDefaultHost(host string) {
	out.DefaultHost = host
}

func (out *DefaultOutputWriter) GetTag() string {
	return out.Tag
}

func New() OutputWriter {
	return NewHEC(utils.Config.Splunk.Hec.Address, utils.Config.Splunk.Hec.Token)
}
