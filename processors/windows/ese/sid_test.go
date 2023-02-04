package ese_test

import (
	"testing"

	"github.com/jurelou/forensibus/processors/windows/ese"
)

func TestValidSid(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"01030000000000055a0000000000000002000000", "S-5-90-0-2"},
		{"010100000000000513000000", "S-5-19"},
		{"010500000000000515000000d4b63e600105027bffd77670ec050000", "S-5-21-1614722772-2063729921-1886836735-1516"},
		{"invalidsid.......", "invalidsid......."},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			res := ese.FormatSID(tt.in)
			if res != tt.out {
				t.Errorf("Invalid SID conversion {%s} should give {%s}, got %s", tt.in, tt.out, res)
			}
		})
	}
}
