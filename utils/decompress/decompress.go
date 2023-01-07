package decompress

import (
	"errors"
	"fmt"
	"os"
	_ "path"
	_ "path/filepath"

	"github.com/h2non/filetype"

	"github.com/jurelou/forensibus/utils"
)

func Decompress(in string, out string) error {
	if exists := utils.FileExists(in); !exists {
		return errors.New(fmt.Sprintf("File %s does not exists", in))
	}
	if err := os.Mkdir(out, 0755); err != nil {
		return err
	}

	// Open a file descriptor
	file, errOpen := os.Open(in)
	if errOpen != nil {
		return errOpen
	}
	defer file.Close()
	head := make([]byte, 261)
	file.Read(head)

	if filetype.IsMIME(head, "application/x-7z-compressed") {
		return DecompressSevenZip(in, out)
		// } else if filetype.IsMIME(head, "application/x-tar") {
		// 	return DecompressSevenZip(in, out)
		// } else if filetype.IsMIME(head, "application/x-xz") {
		// 	return DecompressSevenZip(in, out)
		// } else if filetype.IsMIME(head, "application/vnd.rar") {
		// 	return DecompressSevenZip(in, out)
		// } else if filetype.IsMIME(head, "application/x-bzip2") {
		// 	return DecompressSevenZip(in, out)
		// } else if filetype.IsMIME(head, "application/x-tar") {
		// 	return DecompressSevenZip(in, out)
	} else if filetype.IsMIME(head, "application/x-zip") || filetype.IsMIME(head, "custom/zip") {
		return DecompressZip(in, out)
	}

	// handle := magic.NewMagicHandle(magic.MAGIC_NONE)

	return errors.New(fmt.Sprintf("Decompressing `%s` is not implemented", in))
}
