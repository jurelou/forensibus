package ntfs

import (
	"context"
	"encoding/json"
	"os"

	"www.velocidex.com/golang/go-ntfs/parser"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

type MFTProcessor struct {
	processors.Default
}

func (MFTProcessor) Run(in string, _ *processors.Config, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}
	utils.Log.Infof("Running MFT processor against %s", in)
	fd, err := os.Open(in)
	if err != nil {
		errors.Add(err)
		return errors
	}
	reader, _ := parser.NewPagedReader(fd, 1024, 10000)
	st, err := fd.Stat()
	if err != nil {
		errors.Add(err)
		return errors
	}

	for item := range parser.ParseMFTFile(context.Background(),
		reader, st.Size(), 0x1000, 0x400) {
		serialized, err := json.Marshal(item)
		if err != nil {
			errors.Add(err)
			continue
		}
		e := writer.NewEvent(string(serialized))
		out.WriteEvent(e)
	}
	return errors
}

func init() {
	processors.Register("mft", &MFTProcessor{})
}
