package processors

import (
	"fmt"
	"github.com/jurelou/forensibus/utils"
)

var (
	Registry = make(map[string]Processor)
)

// GetProcessor returns the Component specified by name from `Registry`.
func Get(procName string) (Processor, error) {
	// check if exists
	if _, ok := Registry[procName]; ok {
		return Registry[procName], nil
	}
	return nil, fmt.Errorf("%s is not a registered processor", procName)
}

// Register is called by the `init` function of every `Component` to add
// the caller to the global `Registry` map. If the caller attempts to
// add a `Component` to the registry using the same name as a prior
// `Component` the call will log an error and exit.
func Register(procName string, proc Processor) {

	// check for name collision before adding
	if _, ok := Registry[procName]; ok {
		utils.Log.Warnf("Component: %s has already been added to the registry", procName)
	}
	Registry[procName] = proc
	utils.Log.Debugf("Registered new processor: %s", procName)

}
