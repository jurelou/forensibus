package processors

import (
	"fmt"
	"github.com/pterm/pterm"
	"strings"
	"time"

	"github.com/jurelou/forensibus/utils/processors"
)

func ListProcessors() {
	fmt.Println("Enabled processors:")

	for procName := range processors.Registry {
		fmt.Println("\t", procName)
	}
}

var fakeInstallList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
	"pseudo-dops pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")

func RunSingleProcessor(procName string) {
	fmt.Println("runn", procName)
	pterm.EnableDebugMessages()

	pterm.Debug.Println("Hello, World!")                                                // Print Debug.
	pterm.Info.Println("Hello, World!")                                                 // Print Info.
	pterm.Success.Println("Hello, World!")                                              // Print Success.
	pterm.Warning.Println("Hello, World!")                                              // Print Warning.
	pterm.Error.Println("Errors show the filename and linenumber inside the terminal!") // Print Error.
	pterm.Info.WithShowLineNumber().Println("Other PrefixPrinters can do that too!")    // Print Error.
	// Temporarily set Fatal to false, so that the CI won't crash.
	pterm.Fatal.WithFatal(false).Println("Hello, World!") // Print Fatal.

	p, _ := pterm.DefaultProgressbar.WithTotal(len(fakeInstallList)).WithTitle("Downloading stuff").Start()

	for i := 0; i < p.Total; i++ {
		p.UpdateTitle("Downloading " + fakeInstallList[i])         // Update the title of the progressbar.
		pterm.Success.Println("Downloading " + fakeInstallList[i]) // If a progressbar is running, each print will be printed above the progressbar.
		p.Increment()                                              // Increment the progressbar by one. Use Add(x int) to increment by a custom amount.
		time.Sleep(time.Millisecond * 350)                         // Sleep 350 milliseconds.
	}

}
