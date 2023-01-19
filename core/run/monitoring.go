package run

import (
	// "fmt"
	"time"

	"github.com/pterm/pterm"
)

func MonitorResults(stepsCount int, chans JobChannels, finish chan<- bool) {
	var currentBar *pterm.ProgressbarPrinter
	processes := make(map[string]*CurrentProcess)
	var curProcess string

	for {
		select {
		case j, ok := <-chans.JobResults:
			if !ok {
				chans.JobResults = nil
				break
			}
			key, exists := processes[j.Job.Name]
			if !exists {
				// This should never happens since a CurrentProcess will always be created before a process finishes
				break
			}
			if curProcess == "" {
				// Monitoring just started, set the first progress bar to the current job
				curProcess = j.Job.Name
				currentBar, _ = pterm.DefaultProgressbar.WithElapsedTimeRoundingFactor(time.Millisecond).WithTotal(key.StepsCount).WithTitle(key.ProcessName).WithShowElapsedTime(true).Start()
			}
			key.TerminatedSteps++
			if key.ProcessName == curProcess {
				// Update the progress bar is the process that finished coressponds to the current progress bar
				currentBar.Increment()
			}

			if key.TerminatedSteps >= key.StepsCount {
				// The process has finished, remove it from the map
				delete(processes, key.ProcessName)
				if key.ProcessName != curProcess {
					break
				}
				// If the current process has finished, restart a new progress bar
				currentBar.Stop()
				curProcess = ""
				for k := range processes {
					// The first progress bar is the first map entry
					curProcess = k
					currentBar, _ = pterm.DefaultProgressbar.WithTotal(processes[k].StepsCount).WithTitle(processes[k].ProcessName).WithShowElapsedTime(true).Start()
					break
				}
			}

		case c, ok := <-chans.CurrentProcess: 
			if !ok {
				chans.CurrentProcess = nil
				break
			}
			procName := c.ProcessName
			_, exists := processes[procName]
			if exists {
				processes[procName].StepsCount +=  c.StepsCount
				// Process already exists in the map.
				// This might happen if a processor is called multiple times in the same pipeline
				// Generate a suffix (foobar_1, foobar_2, ...)

				// for i := 1; i < 42; i++ {
				// 	procName = fmt.Sprintf("%s_%d", procName, i)
				// 	_, exists := processes[procName]
				// 	if !exists {
				// 		break
				// 	}
				// }
			}
			if &c != nil {
				processes[procName] = &c
			}

		}

		if chans.JobResults == nil && chans.CurrentProcess == nil {
			break
		}

	}

	finish <- true
}
