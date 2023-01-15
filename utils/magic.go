package utils

import (
	"fmt"

	"github.com/Velocidex/go-magic/magic"
	"github.com/Velocidex/go-magic/magic_files"
)

var handle *magic.Magic

func GetMagic(filePath string) string {
	return handle.File(filePath)
}

func init() {
	handle = magic.NewMagicHandle(magic.MAGIC_NONE)
	// defer handle.Close()
	// handle.SetFlags(magic.MAGIC_MIME_TYPE)
	handle.SetFlags(magic.MAGIC_PRESERVE_ATIME)

	if err := magic_files.LoadDefaultMagic(handle); err != nil {
		fmt.Printf("Could not load magic files: %s", err.Error())
	}
}
