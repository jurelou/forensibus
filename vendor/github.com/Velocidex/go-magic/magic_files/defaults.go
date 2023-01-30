package magic_files

import (
	"github.com/Velocidex/go-magic/magic"
)

func LoadDefaultMagic(handle *magic.Magic) error {
	Init()

	data, err := ReadFile("magic_files/magic_rules.mgc")
	if err != nil {
		return err
	}
	handle.LoadBuffer(string(data))

	return nil
}
