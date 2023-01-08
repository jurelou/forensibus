package prefetch

// References from
// https://github.com/libyal/libscca/blob/master/documentation/Windows%20Prefetch%20File%20(PF)%20format.asciidoc

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

func LoadPrefetch(reader io.ReaderAt) (*PrefetchInfo, error) {
	profile := NewPrefetchProfile()
	header := profile.MAMHeader(reader, 0)
	if header.Signature() == "MAM\x04" {
		// Need to decompress it in memory.
		data := make([]byte, header.UncompressedSize())
		n, err := reader.ReadAt(data, int64(header.Size()))
		if err != nil && err != io.EOF {
			return nil, err
		}

		decompressed, err := LZXpressHuffmanDecompressWithFallback(
			data[:n], int(header.UncompressedSize()))
		if err != nil {
			return nil, err
		}

		if Prefetch_debug {
			fd, err := os.OpenFile(
				"/tmp/prefet_dump.bin", os.O_RDWR|os.O_CREATE, 0755)
			if err != nil {
				return nil, err
			}
			fd.Write(decompressed)
			fd.Close()
		}

		reader = bytes.NewReader(decompressed)
	}

	scca_header := profile.SCCAHeader(reader, 0)

	if scca_header.Signature() != "SCCA" {
		return nil, errors.New("Not a prefetch file (bad signature).")
	}

	self := &PrefetchInfo{
		Executable: scca_header.Executable(),
		FileSize:   scca_header.FileSize(),
		Hash:       fmt.Sprintf("0x%08X", scca_header.Hash()),
		Version:    scca_header.Version().Name,
	}

	switch self.Version {

	case "WinXP":
		file_info := profile.FileInformationXP(
			scca_header.Reader,
			scca_header.Offset+int64(scca_header.Size()))

		self.LastRunTimes = append(self.LastRunTimes, file_info.LastRunTime().Time)
		self.FilesAccessed = file_info.Filenames()
		self.RunCount = file_info.RunCount()

	case "Vista":
		file_info := profile.FileInformationVista(
			scca_header.Reader,
			scca_header.Offset+int64(scca_header.Size()))

		self.LastRunTimes = append(self.LastRunTimes, file_info.LastRunTime().Time)
		self.FilesAccessed = file_info.Filenames()
		self.RunCount = file_info.RunCount()

	case "Win10", "Win8.1":
		file_info := profile.FileInformationWin10(
			scca_header.Reader,
			scca_header.Offset+int64(scca_header.Size()))

		for _, last_run := range file_info.LastRunTimes() {
			if last_run.After(time.Now()) {
				continue
			}
			self.LastRunTimes = append(self.LastRunTimes, last_run.Time)
		}

		self.FilesAccessed = file_info.Filenames()
		self.RunCount = file_info.RunCount1()
		if self.RunCount == 0 {
			self.RunCount = file_info.RunCount2()
		}
	}

	return self, nil
}
