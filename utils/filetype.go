package utils

import (
	// "fmt"
	"github.com/h2non/filetype"
)

var zipType = filetype.NewType("Archive", "custom/zip")

func zipMatcher(buf []byte) bool {
	// Matches PK\3\4
	// TODO: this is not correct, zip file header is not forced to be at the begining of the file
	return len(buf) > 3 && buf[0] == 0x50 && buf[1] == 0x4b && buf[2] == 0x03 && buf[3] == 0x04
}

func init() {
	// fmt.Println("hello")
	filetype.AddMatcher(zipType, zipMatcher)
}
