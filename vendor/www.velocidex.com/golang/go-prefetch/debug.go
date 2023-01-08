package prefetch

import "fmt"

var (
	Prefetch_debug = false
)

func Printf(fmt_str string, args ...interface{}) {
	if Prefetch_debug {
		fmt.Printf(fmt_str, args...)
	}
}
