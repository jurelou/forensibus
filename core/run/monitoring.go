package run

import (
	"time"

	"github.com/pterm/pterm"

	"github.com/jurelou/forensibus/utils"
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
			if utils.IsErrored(j.Status) {
				pterm.Error.Printfln("Processor %s failed (%d errors) against %s: %v", j.Job.Name, len(j.Errors), j.Job.Step.NextArtifact, j.Errors)
			}
			key, exists := processes[j.Job.ProcessId]
			if !exists {
				pterm.Error.Println("Undefined job")

				// This should never happens since a CurrentProcess will always be created before a process finishes
				break
			}
			if curProcess == "" {
				// Monitoring just started, set the first progress bar to the current job
				curProcess = j.Job.ProcessId
				currentBar, _ = pterm.DefaultProgressbar.WithElapsedTimeRoundingFactor(time.Millisecond).WithTotal(key.StepsCount).WithTitle(key.ProcessName).WithShowElapsedTime(true).Start()
			}
			key.TerminatedSteps++

			if key.ProcessId == curProcess {
				// Update the progress bar is the process that finished coressponds to the current progress bar
				currentBar.Increment()
			}

			if key.TerminatedSteps >= key.StepsCount {
				// The process has finished, remove it from the map
				delete(processes, j.Job.ProcessId)
				if key.ProcessId != curProcess {
					break
				}
				// If the current process has finished, restart a new progress bar
				currentBar.Stop()
				curProcess = ""
				for k := range processes {
					// The first progress bar is the first map entry
					curProcess = k
					currentBar, _ = pterm.DefaultProgressbar.WithTotal(processes[k].StepsCount).WithTitle(processes[k].ProcessName).WithShowElapsedTime(true).Start()
					currentBar.Add(processes[k].TerminatedSteps)
					break
				}
			}

		case c, ok := <-chans.CurrentProcess:
			if !ok {
				chans.CurrentProcess = nil
				break
			}
			_, exists := processes[c.ProcessId]
			if exists {
				// This should never happen ...
				pterm.Error.Println("Duplicate process identifier detected !!!!!")
				break
			}
			if &c != nil {
				processes[c.ProcessId] = &c
			}

		}

		if chans.JobResults == nil && chans.CurrentProcess == nil {
			break
		}

	}

	finish <- true
}
