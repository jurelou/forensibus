package processors_test

import (
	"fmt"
	"testing"

	"github.com/jurelou/forensibus/utils/processors"
)

func TestGetEvtx(t *testing.T) {
	_, err := processors.Get("evtx")
	if err == nil {
		t.Errorf("Processor `this_should_not_exists` should not exists")
	}
}

func TestGetInvalid(t *testing.T) {
	fmt.Println("ok")
	_, err := processors.Get("this_should_not_exists")
	if err == nil {
		t.Errorf("Processor `this_should_not_exists` should not exists")
	}
}
