package ntfs_test

import (
	"testing"

	"github.com/jurelou/forensibus/processors/windows/ntfs"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

type MonkeyWriter struct {
	writer.DefaultOutputWriter
	cb func(*writer.Event)
}

func NewMonkeyWriter(cb func(*writer.Event)) *MonkeyWriter {
	return &MonkeyWriter{cb: cb}
}

func (w MonkeyWriter) WriteEvent(event *writer.Event) {
	w.cb(event)
}

func TestMftInvalidFile(t *testing.T) {
	calls := 0
	cb := func(e *writer.Event) {
		calls++
	}
	writer := NewMonkeyWriter(cb)
	proc := ntfs.MFTProcessor{}
	errs := proc.Run("/tmp/this/file/does/not/exists", &processors.Config{RawConfig: map[string]string{}}, writer)

	if calls != 0 {
		t.Errorf("MFT parser returned too many results for invalid_file")
	}

	if errs.Empty() {
		t.Errorf("MFT should fail on invalid file")
	}
	if errs.AsStrings()[0] != "open /tmp/this/file/does/not/exists: no such file or directory" {
		t.Errorf("Invalid error: %s", errs.AsStrings()[0])
	}
}

func TestMft(t *testing.T) {
	calls := 0
	wants := `{"EntryNumber":47,"InUse":true,"ParentEntryNumber":39,"FullPath":"\u003cunknown\u003e/time_for_a_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super__super_super_super_super_super_super_super_super_longname.txt","FileName":"time_for_a_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super_super__super_super_super_super_super_super_super_super_longname.txt","FileSize":31,"ReferenceCount":1,"IsDir":false,"Created0x10":"2017-04-19T17:39:37.5419077-07:00","Created0x30":"2017-04-19T17:39:37.5419077-07:00","LastModified0x10":"2017-04-19T17:40:33.7241746-07:00","LastModified0x30":"2017-04-19T17:39:37.5419077-07:00","LastRecordChange0x10":"2017-04-19T17:40:33.7241746-07:00","LastRecordChange0x30":"2017-04-19T17:40:05.1183341-07:00","LastAccess0x10":"2017-04-19T17:39:37.5419077-07:00","LastAccess0x30":"2017-04-19T17:39:37.5419077-07:00"}`
	cb := func(e *writer.Event) {
		if e.Event != wants {
			t.Errorf("Invalid MFT entry %s", e.Event)
		}
		calls++
	}
	writer := NewMonkeyWriter(cb)
	proc := ntfs.MFTProcessor{}
	errs := proc.Run("../../../datasets/artifacts/mft/entry_super_long_name_001", &processors.Config{RawConfig: map[string]string{}}, writer)

	if !errs.Empty() {
		t.Errorf("MFT had some errors: %v", errs.AsStrings())
	}
	if calls != 1 {
		t.Errorf("MFT parser returned too many results for entry_super_long_name_001")
	}
}
