package appcompatcache

import (
	"bytes"
	"time"
)

func filetimeToUnixtime(ft uint64) uint64 {
	return (ft - 11644473600000*10000) / 10000000
}

func (self *Win10CreatorsEntry) Path() string {
	return ParseUTF16String(
		self.Reader, self.Offset+self.Profile.Off_Win10CreatorsEntry_Path__,
		int64(self.PathSize()))
}

func (self *Win10CreatorsEntry) IsValidSignature() bool {
	switch self.Signature() {
	case 0x73743031, 0x73743030:
		return true
	}

	return false
}

func (self *Win10CreatorsEntry) LastMod(header *Win10CreatorsHeader) uint64 {
	offset := self.Offset + self.Profile.Off_Win10CreatorsEntry_Path__ +
		int64(self.PathSize())

	if self.Signature() == 0x73743030 {
		offset += 10
	} else if header.HeaderSize() == 0x80 {
		offset += 10
	}

	return filetimeToUnixtime(ParseUint64(self.Reader, offset))
}

func (self *Win10CreatorsEntry) Next() *Win10CreatorsEntry {
	return self.Profile.Win10CreatorsEntry(self.Reader,
		self.Offset+self.Profile.Off_Win10CreatorsEntry_PathSize+
			int64(self.DataSize()))
}

func ParseValueData(buffer []byte) []*CacheEntry {
	result := []*CacheEntry{}

	profile := NewAppCompatibilityProfile()
	fd := bytes.NewReader(buffer)

	header := profile.Win10CreatorsHeader(fd, 0)
	offset := int64(header.HeaderSize())
	if offset == 0 {
		offset = 0x80
	}

	for entry := profile.Win10CreatorsEntry(fd, offset); entry.IsValidSignature(); entry = entry.Next() {
		ts := int64(entry.LastMod(header))
		if ts < 0 || ts > 2000000000 {
			ts = 0
		}

		result = append(result, &CacheEntry{
			Name:  entry.Path(),
			Epoch: entry.LastMod(header),
			Time:  time.Unix(ts, 0),
		})
	}

	return result
}
