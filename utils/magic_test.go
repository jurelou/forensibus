package utils_test

import (
	"strings"
	"testing"

	"github.com/jurelou/forensibus/utils"
)

func TestMagic(t *testing.T) {
	magic := utils.GetMagic("../datasets/archives/Simple.7z")
	if magic != "application/x-7z-compressed" {
		t.Errorf("Invalid magic for Simple.7z")
	}
}

func TestInvalidMagicFile(t *testing.T) {
	magic := utils.GetMagic("/this/does/not/exists")
	if !strings.Contains(magic, "No such file or directory") {
		t.Errorf("Invalid error code from unknown file")
	}
}
