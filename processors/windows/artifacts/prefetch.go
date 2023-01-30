package windows_artifacts

import (
	"encoding/json"
	"os"
	"time"

	"www.velocidex.com/golang/go-prefetch"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

type PrefetchEntry struct {
	Executable    string      `json:"ExecutableName"`
	FileSize      uint32      `json:"FileSize"`
	Hash          string      `json:"Hash"`
	Version       string      `json:"Version"`
	LastRunTimes  []time.Time `json:"LastRunTimes"`
	FilesAccessed []string    `json:"FilesAccessed"`
	RunCount      uint32      `json:"RunCount"`
}

type PrefetchProcessor struct {
	processors.Default
}

func (PrefetchProcessor) parsePrefetch(in string) (PrefetchEntry, error) {
	fd, err := os.Open(in)
	if err != nil {
		utils.Log.Warnf("Could not open prefetch file `%s`", in)
		return PrefetchEntry{}, err
	}
	pf, err := prefetch.LoadPrefetch(fd)
	if err != nil {
		utils.Log.Warnf("Prefetch file `%s` is invalid: `%s`", in, err.Error())
		return PrefetchEntry{}, err
	}
	return PrefetchEntry(*pf), nil
}

func (proc PrefetchProcessor) Run(in string, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}

	entry, err := proc.parsePrefetch(in)
	if err != nil {
		errors.Add(err)
		return errors
	}

	json, err := json.Marshal(entry)
	if err != nil {
		errors.Add(err)
		return errors
	}

	e := writer.NewEvent(string(json))
	out.WriteEvent(e)
	return errors
}

func init() {
	processors.Register("prefetch", PrefetchProcessor{})
}
