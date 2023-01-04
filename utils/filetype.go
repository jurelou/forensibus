package utils

import (
	// "fmt"
	"github.com/h2non/filetype"

)

var fooType = filetype.NewType("foo", "foo/foo")

func fooMatcher(buf []byte) bool {
  return len(buf) > 1 && buf[0] == 0x01 && buf[1] == 0x02
}

func init() {
	// fmt.Println("hello")
	filetype.AddMatcher(fooType, fooMatcher)
}