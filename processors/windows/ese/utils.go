package windows_ese

import (
	"os"
	"fmt"
	"io"
	"math/bits"
	"encoding/hex"
	"encoding/binary"
	"www.velocidex.com/golang/go-ese/parser"
	// "github.com/Velocidex/ordereddict"

	// "github.com/jurelou/forensibus/utils/processors"
	// "github.com/jurelou/forensibus/utils/writer"
)

func GetCatalog(fd *os.File) (*parser.Catalog, error) {
	ese_ctx, err := parser.NewESEContext(fd)
	if err != nil {
		return nil, fmt.Errorf("unable to parse ESE database %s: %w", fd.Name(), err)
	}
	catalog, err := parser.ReadCatalog(ese_ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to parse ESE catalog %s: %w", fd.Name(), err)
	}
	return catalog, nil

}

type MiscProfile struct {
    Off_Misc_Misc int64
    Off_SID_Revision int64
    Off_SID_SubAuthCount int64
    Off_SID_Authority int64
    Off_SID_Authority2 int64
    Off_SID_Subauthority int64
}

type SID struct {
    Reader io.ReaderAt
    Offset int64
    Profile *MiscProfile
}
func (self *SID) Revision() byte {
	return ParseUint8(self.Reader, self.Profile.Off_SID_Revision + self.Offset)
 }
 
 func (self *SID) SubAuthCount() byte {
	return ParseUint8(self.Reader, self.Profile.Off_SID_SubAuthCount + self.Offset)
 }
 
 func (self *SID) Authority() uint16 {
	return ParseUint16(self.Reader, self.Profile.Off_SID_Authority + self.Offset)
 }
 
 func (self *SID) Authority2() uint32 {
	return ParseUint32(self.Reader, self.Profile.Off_SID_Authority2 + self.Offset)
 }
 
 func (self *SID) Subauthority() []uint32 {
	return ParseArray_uint32(self.Profile, self.Reader, self.Profile.Off_SID_Subauthority + self.Offset, 100)
 }

func (self *SID) String() string {
	result := fmt.Sprintf("S-%d", uint64(bits.ReverseBytes16(self.Authority()))<<32+
		uint64(bits.ReverseBytes32(self.Authority2())))

	sub_authorities := self.Subauthority()
	for i := 0; i < int(self.SubAuthCount()); i++ {
		if i > len(sub_authorities) {
			break
		}

		sub := sub_authorities[i]
		result += fmt.Sprintf("-%d", sub)
	}
	return result
}

func ParseArray_uint32(profile *MiscProfile, reader io.ReaderAt, offset int64, count int) []uint32 {
    result := make([]uint32, 0, count)
    for i:=0; i<count; i++ {
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
func NewMiscProfile() *MiscProfile {
    self := &MiscProfile{0,0,1,2,4,8}
    return self
}

func (self *MiscProfile) SID(reader io.ReaderAt, offset int64) *SID {
    return &SID{Reader: reader, Offset: offset, Profile: self}
}

func FormatGUID(hexencoded string) []byte {
	if len(hexencoded) == 0 {
		return nil
	}

	buffer, err := hex.DecodeString(hexencoded)
	if err != nil {
		return nil
	}
	return buffer
	// profile := NewMiscProfile()
	// return profile.SID(&BufferReaderAt{Buffer: buffer}, 0).String()
}



type BufferReaderAt struct {
	Buffer []byte
}

func (self *BufferReaderAt) ReadAt(buf []byte, offset int64) (int, error) {
	to_read := int64(len(buf))
	if offset < 0 {
		to_read += offset
		offset = 0
	}

	if offset+to_read > int64(len(self.Buffer)) {
		to_read = int64(len(self.Buffer)) - offset
	}

	if to_read < 0 {
		return 0, nil
	}

	n := copy(buf, self.Buffer[offset:offset+to_read])

	return n, nil
}