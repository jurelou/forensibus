package common_processors_test

import (
	"testing"

	common_processors "github.com/jurelou/forensibus/processors"
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

type TestCase struct {
	filePath            string
	expectedFirstOutput string
	expectedRowsCount   int
	shouldFail          bool
}

func TestIni(t *testing.T) {
	tests := []TestCase{
		{
			filePath:            "../datasets/artifacts/evtx/sysmon_3.evtx",
			expectedFirstOutput: "",
			expectedRowsCount:   0,
			shouldFail:          true,
		},
		{
			filePath:            "/tmp/this/does/not/exists",
			expectedFirstOutput: "",
			expectedRowsCount:   0,
			shouldFail:          true,
		},
		{
			filePath:            "../datasets/artifacts/ini/ads",
			expectedFirstOutput: `{"ZoneTransfer":{"HostUrl":"https://example.com","ZoneId":"3"},"some":"test"}`,
			expectedRowsCount:   1,
			shouldFail:          false,
		},
	}

	proc := common_processors.IniProcessor{}
	for _, tt := range tests {

		calls := 0
		cb := func(e *writer.Event) {
			if tt.expectedFirstOutput != "" && tt.expectedFirstOutput != e.Event {
				t.Errorf("Invalid output from %s: %s", tt.filePath, e.Event)
			}
			calls++
		}
		writer := NewMonkeyWriter(cb)

		t.Run(tt.filePath, func(t *testing.T) {
			errs := proc.Run(tt.filePath, &processors.Config{RawConfig: map[string]string{}}, writer)
			if !tt.shouldFail && !errs.Empty() {
				t.Errorf("File `%s` returned an error: %v", tt.filePath, errs)
			}
		})

		if calls != tt.expectedRowsCount {
			t.Errorf("File `%s` returned %d rows instead of %d", tt.filePath, calls, tt.expectedRowsCount)
		}
	}
}
