package utils_test

import (
	"testing"

	"github.com/jurelou/forensibus/utils"
)

func TestConfig(t *testing.T) {
	err := utils.Configure()
	if err != nil {
		t.Errorf("Error while Loading config %s", err.Error())
	}
	err = utils.ReloadConfig()
	if err != nil {
		t.Errorf("Error while Reloading config %s", err.Error())
	}
	passwds := utils.Config.ArchivePasswords
}
