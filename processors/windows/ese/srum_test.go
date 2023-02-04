package ese_test

import (
	"testing"

	"github.com/jurelou/forensibus/processors/windows/ese"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

type MonkeyWriter struct {
	writer.DefaultOutputWriter
	cb func(*writer.Event)
}

func NewMonkeyWriter(cb func(*writer.Event)) *MonkeyWriter {
	return &MonkeyWriter{cb: cb}
}

func (w MonkeyWriter) WriteEvent(event *writer.Event) {
	w.cb(event)
}

func TestSRUM(_ *testing.T) {
	cb := func(e *writer.Event) {
		// appId, exists := e.Event.GetString("AppId")
		// if !exists {
		// 	t.Errorf("Srum event does not have an `AppId` field")
		// }
	}
	writer := NewMonkeyWriter(cb)
	proc := ese.SrumProcessor{}
	proc.Run("../../../datasets/artifacts/ese/srum.data", &processors.Config{RawConfig: map[string]string{}}, writer)
}
