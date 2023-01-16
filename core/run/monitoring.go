package run

import (
	"github.com/pterm/pterm"
	// "github.com/jurelou/forensibus/utils"
)

func MonitorResults(stepsCount int, chans JobChannels, finish chan<- bool) {
	// globalBar, _ := pterm.DefaultProgressbar.WithTotal(stepsCount).WithTitle("Total").WithRemoveWhenDone(false).Start()

	var currentBar *pterm.ProgressbarPrinter
	var current *CurrentProcess

	for {
		select {
		case _, ok := <-chans.JobResults:
			if !ok {
				chans.JobResults = nil
				break
			}
			currentBar.Increment()
			// fmt.Println("new result", a, ok)
		case c, ok := <-chans.CurrentProcess:
			if current != nil {
				// It's not the first processor, it means the last processor ended
				// Print out that the last processor finished
				pterm.Success.Printfln("Terminated %s processor", current.ProcessName)
				// Increment the global progress bar
				// globalBar.Increment()
				// Close the current progress bar
				currentBar.Stop()
			}
			// Start a new progress bar
			currentBar, _ = pterm.DefaultProgressbar.WithTotal(c.StepsCount).WithTitle(c.ProcessName).WithRemoveWhenDone(true).Start()

			current = &c
			if !ok {
				chans.CurrentProcess = nil
				break
			}
		}

		if chans.JobResults == nil && chans.CurrentProcess == nil {
			break
		}

	}

	finish <- true
}
