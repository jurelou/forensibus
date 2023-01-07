package utils

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	err := Configure()
	if err != nil {
		t.Errorf("Error while Loading config %s", err.Error())
	}
	err = Reload()
	if err != nil {
		t.Errorf("Error while Reloading config %s", err.Error())
	}
	passwds := Config.ArchivePasswords
	fmt.Println(passwds)
}
