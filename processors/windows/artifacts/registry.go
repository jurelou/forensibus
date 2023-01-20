package windows_artifacts

import (
	"os"
	"fmt"
	// "time"

	"www.velocidex.com/golang/regparser"
	// "www.velocidex.com/golang/regparser/appcompatcache"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

const (
	appcompatcache_path = "/ControlSet001/Control/Session Manager/AppCompatCache"
)

type RegistryProcessor struct {
}

func (proc RegistryProcessor) Configure() error {
	return nil
}

func (proc RegistryProcessor) Run(in string, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}
	fd, err := os.Open(in)
	if err != nil {
		utils.Log.Warnf("Could not open registry hive `%s`", in)
		errors.Add(err)
		return errors
	}

	fmt.Println(">>>>>>>>>>>>", in)
	registry, err := regparser.NewRegistry(fd)
	if err != nil {
		errors.Add(err)
		return errors
	}
	key := registry.OpenKey(appcompatcache_path)
	if key != nil {
		fmt.Println("<<<<<<<<<<<", key)

	}
	// time.Sleep(1 * time.Second)

	return errors
}

func init() {
	processors.Register("registry", RegistryProcessor{})
}
