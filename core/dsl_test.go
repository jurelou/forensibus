package core_test

import (
	// "fmt"
	// "reflect"
	"strings"
	"testing"

	"github.com/jurelou/forensibus/core"
)

func TestInvalidFile(t *testing.T) {
	_, err := core.LoadDSLFile("../datasets/this.does.not.exists")

	if err == nil {
		t.Errorf("FindFiles should fail to find 'this.does.not.exists'")
	}
	if !strings.Contains(err.Error(), "The configuration file ../datasets/this.does.not.exists does not exist") {
		t.Errorf("Invalid error message: %s", err.Error())
	}
}
