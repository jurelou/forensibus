package utils_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jurelou/forensibus/utils"
)

func TestMagic(t *testing.T) {
	magic := utils.GetMagic("../datasets/archives/Simple.7z")
	if !strings.Contains(magic, "7-zip archive data, version") {
		t.Errorf("Invalid magic for Simple.7z")
	}
}

func TestInvalidMagicFile(t *testing.T) {
	magic := utils.GetMagic("/this/does/not/exists")
	if !strings.Contains(magic, "No such file or directory") {
		t.Errorf("Invalid error code from unknown file")
	}
}

func TestEvtxMagic(t *testing.T) {
	magic := utils.GetMagic("../datasets/artifacts/evtx/sysmon_3.evtx")
	fmt.Println(magic)
	// if !strings.Contains(magic, "No such file or directory") {
	// 	t.Errorf("Invalid error code from unknown file")
	// }
}