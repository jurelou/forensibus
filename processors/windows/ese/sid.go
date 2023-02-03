// Code mostly taken from https://github.com/Velocidex/velociraptor/blob/eced6fc858e659281d80749661340ba42e8e099e/vql/parsers/ese/sid.go
package ese

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"math/bits"
)

const (
	// RevisionLevelOffset     int64 = 0
	SubAuthorityCountOffset int64 = 1
	AuthorityOffset         int64 = 2
	Authority2Offset        int64 = 4
	SubauthorityOffset      int64 = 8
)

func FormatSID(hexencoded string) string {
	if len(hexencoded) == 0 {
		return hexencoded
	}

	buffer, err := hex.DecodeString(hexencoded)
	if err != nil {
		return hexencoded
	}
	sid := &SID{Reader: &BufferReaderAt{Buffer: buffer}, Offset: 0}
	return sid.String()
}

type SID struct {
	Reader io.ReaderAt
	Offset int64
}

// func (s *SID) Revision() byte {
// 	return ParseUint8(s.Reader, RevisionLevelOffset+s.Offset)
// }

func (s *SID) SubAuthCount() byte {
	return ParseUint8(s.Reader, SubAuthorityCountOffset+s.Offset)
}

func (s *SID) Authority() uint16 {
	return ParseUint16(s.Reader, AuthorityOffset+s.Offset)
}

func (s *SID) Authority2() uint32 {
	return ParseUint32(s.Reader, Authority2Offset+s.Offset)
}

func (s *SID) Subauthority() []uint32 {
	return ParseArray_uint32(s.Reader, SubauthorityOffset+s.Offset, 100)
}

func (s *SID) String() string {
	result := fmt.Sprintf("S-%d", uint64(bits.ReverseBytes16(s.Authority()))<<32+
		uint64(bits.ReverseBytes32(s.Authority2())))

	sub_authorities := s.Subauthority()
	for i := 0; i < int(s.SubAuthCount()); i++ {
		if i > len(sub_authorities) {
			break
		}

		sub := sub_authorities[i]
		result += fmt.Sprintf("-%d", sub)
	}
	return result
}

func ParseArray_uint32(reader io.ReaderAt, offset int64, count int) []uint32 {
	result := make([]uint32, 0, count)
	for i := 0; i < count; i++ {
		value := ParseUint32(reader, offset)
		result = append(result, value)
		offset += int64(4)
	}
	return result
}

func ParseUint32(reader io.ReaderAt, offset int64) uint32 {
	data := make([]byte, 4)
	_, err := reader.ReadAt(data, offset)
	if err != nil {
		return 0
	}
	return binary.LittleEndian.Uint32(data)
}

func ParseUint8(reader io.ReaderAt, offset int64) byte {
	result := make([]byte, 1)
	_, err := reader.ReadAt(result, offset)
	if err != nil {
		return 0
	}
	return result[0]
}

func ParseUint16(reader io.ReaderAt, offset int64) uint16 {
	data := make([]byte, 2)
	_, err := reader.ReadAt(data, offset)
	if err != nil {
		return 0
	}
	return binary.LittleEndian.Uint16(data)
}

type BufferReaderAt struct {
	Buffer []byte
}

func (buff *BufferReaderAt) ReadAt(buf []byte, offset int64) (int, error) {
	to_read := int64(len(buf))
	if offset < 0 {
		to_read += offset
		offset = 0
	}

	if offset+to_read > int64(len(buff.Buffer)) {
		to_read = int64(len(buff.Buffer)) - offset
	}

	if to_read < 0 {
		return 0, nil
	}

	n := copy(buf, buff.Buffer[offset:offset+to_read])

	return n, nil
}
