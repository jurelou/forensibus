package processors

import (
	"fmt"
	"log"

)

var (
	Registry = make(map[string]Processor)
)

func Toto() {
	fmt.Println("salut")
}

// Register is called by the `init` function of every `Component` to add
// the caller to the global `Registry` map. If the caller attempts to
// add a `Component` to the registry using the same name as a prior
// `Component` the call will log an error and exit.
func Register(procName string, proc Processor) {

	// check for name collision before adding
	if _, ok := Registry[procName]; ok {
		log.Fatalf("Component: %s has already been added to the registry", procName)
	}
	Registry[procName] = proc
	fmt.Println("Registered new processor:", procName)

}