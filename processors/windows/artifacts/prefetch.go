package windows_artifacts

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"www.velocidex.com/golang/go-prefetch"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
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
	Input string `yaml:"input"`
}

func (p PrefetchProcessor) Configure() error {
	return nil
}
func (p PrefetchProcessor) Run(in string) error {
	utils.Log.Debugf("Run pf processor against `%s`", in)

	fd, _ := os.Open(in)
	pf, _ := prefetch.LoadPrefetch(fd)
	entry := PrefetchEntry(*pf)

	json, _ := json.Marshal(entry)
	fmt.Println(string(json))
	return nil
}

func init() {
	processors.Register("prefetch", PrefetchProcessor{})
}
