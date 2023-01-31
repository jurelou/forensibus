package windows_artifacts

import (
	"fmt"
	"os"

	"www.velocidex.com/golang/regparser"

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

func rec(nk *regparser.CM_KEY_NODE) {
	fmt.Println("-- ", nk.LastWriteTime(), nk.Name())
	for _, v := range nk.Values() {
		fmt.Println("   >", v.Name(), v.ValueName(), v.ValueData())
	}

	for _, subkey := range nk.Subkeys() {
		// fmt.Println("<<<<<<<<<<<", subkey.Name())
		fmt.Println("-- ", subkey.LastWriteTime(), subkey.Name())
		return
		rec(subkey)
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

	// fmt.Println(">>>>>>>>>>>>", in)
	registry, err := regparser.NewRegistry(fd)
	// accKey := registry.OpenKey(appcompatcache_path)
	// if accKey != nil {
	// 	fmt.Println("APP COMPAT CACHE HERE", in)
	// 	return errors
	// }
	if err != nil {
		errors.Add(err)
		return errors
	}
	rootCell := registry.Profile.HCELL(registry.Reader, 0x1000+int64(registry.BaseBlock.RootCell()))
	nk := rootCell.KeyNode()
	if nk == nil {
		errors.Add(fmt.Errorf("Could not find root cell from %s", in))
		return errors
	}

	fmt.Println("@@@@@@@@@@@", in)
	rec(nk)

	// time.Sleep(1 * time.Second)

	return errors
}

func init() {
	processors.Register("registry", RegistryProcessor{})
}
