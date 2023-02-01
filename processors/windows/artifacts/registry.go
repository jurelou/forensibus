package windows_artifacts

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"www.velocidex.com/golang/regparser"
	"www.velocidex.com/golang/regparser/appcompatcache"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

const (
	appcompatcache_path = "/ControlSet001/Control/Session Manager/AppCompatCache"
)

type RegistryProcessor struct {
	processors.Default
}

type RegistryEntry struct {
	LastWriteTime string
	Key           string
	KeyName       string
	Value         string
	ValueName     string
	// KeyName string
}

type AppCompatCacheEntry struct {
	Time time.Time
	Name string
	Id   int
}

func dumpAppCompatCache(nk *regparser.CM_KEY_NODE, out writer.OutputWriter) {
	out.SetDefaultSourceType("forensibus_appcompatcache")

	for _, value := range nk.Values() {
		if value.ValueName() != "AppCompatCache" {
			continue
		}
		for idx, entry := range appcompatcache.ParseValueData(value.ValueData().Data) {
			// fmt.Printf("%d: %v  %v\n", idx, entry.Time, entry.Name)
			entry := AppCompatCacheEntry{
				Time: entry.Time,
				Name: entry.Name,
				Id:   idx,
			}
			jsonEntry, _ := json.Marshal(entry)
			e := writer.NewEvent(string(jsonEntry))
			out.WriteEvent(e)
		}
	}
}

func rec(nk *regparser.CM_KEY_NODE, path string, out writer.OutputWriter) {
	keyName := nk.Name()
	lastWrite := nk.LastWriteTime().GoString()
	path = path + "/" + keyName

	for _, v := range nk.Values() {
		valueData := v.ValueData()
		entry := RegistryEntry{
			LastWriteTime: lastWrite,
			Key:           path,
			KeyName:       keyName,
			ValueName:     v.ValueName(),
		}

		if valueData.String != "" {
			entry.Value = valueData.String
		} else if len(valueData.MultiSz) > 0 {
			entry.Value = strings.Join(valueData.MultiSz, ",")
		} else if valueData.Uint64 != 0 {
			entry.Value = strconv.FormatUint(valueData.Uint64, 10)
		} else {
			entry.Value = fmt.Sprintf("<omit %s>", v.TypeString())
		}
		jsonEntry, _ := json.Marshal(entry)
		e := writer.NewEvent(string(jsonEntry))
		out.WriteEvent(e)
	}

	for _, subkey := range nk.Subkeys() {
		rec(subkey, path, out)
	}
}

func (RegistryProcessor) Run(in string, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}
	fd, err := os.Open(in)
	if err != nil {
		utils.Log.Warnf("Could not open registry hive `%s`", in)
		errors.Add(err)
		return errors
	}
	registry, err := regparser.NewRegistry(fd)
	if err != nil {
		errors.Add(err)
		return errors
	}

	accKey := registry.OpenKey(appcompatcache_path)
	if accKey != nil {
		dumpAppCompatCache(accKey, out)
	}

	rootCell := registry.Profile.HCELL(registry.Reader, 0x1000+int64(registry.BaseBlock.RootCell()))
	nk := rootCell.KeyNode()
	if nk == nil {
		errors.Add(fmt.Errorf("Could not find root cell from %s", in))
		return errors
	}
	out.SetDefaultSourceType("forensibus_registry")
	rec(nk, "", out)

	return errors
}

func init() {
	processors.Register("registry", RegistryProcessor{})
}
