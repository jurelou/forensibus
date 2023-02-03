package windows_ese

import (
    "io"
	"bytes"
	"encoding/hex"
	"unicode/utf8"
	"unicode/utf16"
    "encoding/binary"
)

func UTF16BytesToUTF8(b []byte, o binary.ByteOrder) string {
	if len(b) < 2 {
		return ""
	}

	if b[0] == 0xff && b[1] == 0xfe {
		o = binary.BigEndian
		b = b[2:]
	} else if b[0] == 0xfe && b[1] == 0xff {
		o = binary.LittleEndian
		b = b[2:]
	}

	utf := make([]uint16, (len(b)+(2-1))/2)

	for i := 0; i+(2-1) < len(b); i += 2 {
		utf[i/2] = o.Uint16(b[i:])
	}
	if len(b)/2 < len(utf) {
		utf[len(utf)-1] = utf8.RuneError
	}

	return string(utf16.Decode(utf))
}


func FormatUtf16String(hexencoded string) string {
	buffer, err := hex.DecodeString(hexencoded)
	if err != nil {
		return hexencoded
	}
	reader := BufferReaderAt{Buffer: buffer}
	data := make([]byte, 1024)

	n, err := reader.ReadAt(data, 0)
	if err != nil && err != io.EOF {
	  return ""
	}
 
	idx := bytes.Index(data[:n], []byte{0, 0})
	if idx < 0 {
	   idx = n-1
	}
	return UTF16BytesToUTF8(data[0:idx+1], binary.LittleEndian)
}
