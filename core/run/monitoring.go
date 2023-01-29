package run

import (
	"time"

	"github.com/pterm/pterm"
	"github.com/jurelou/forensibus/utils"
)

func MonitorResults(done chan<- bool, startedProcesses <-chan ProcessStarted, endedTasks <-chan TaskEnded) {
	var currentProcess string
	var pBar *pterm.ProgressbarPrinter
	processes := make(map[string]*ProcessStarted)

	for {
		select {
		case proc, ok := <-startedProcesses:
			if !ok {
				startedProcesses = nil
				break
			}
			_, exists := processes[proc.ProcessId]
			if exists {
				// This should never happen ... (ProcessId is a valid UUID4)
				pterm.Error.Println("Duplicate process ID detected . . .")
				break
			}
			if &proc != nil {
				processes[proc.ProcessId] = &proc
			}

			if currentProcess == "" {
				// It is the first processor 
				currentProcess = proc.ProcessId
				pBar, _ = pterm.DefaultProgressbar.WithElapsedTimeRoundingFactor(time.Millisecond).WithTotal(proc.TasksCount).WithTitle(proc.Name).WithShowElapsedTime(true).Start()
			}

		case task, ok := <-endedTasks:
			if !ok {
				endedTasks = nil
				break
			}
			if utils.IsErrored(task.Status) {
				pterm.Error.Printfln("Processor %s failed (%d errors) against %s: %v", task.Name, len(task.Errors), task.Source, task.Errors)
			}
			key, exists := processes[task.ProcessId]
			if !exists {
				pterm.Error.Println("Undefined job . . .")
				break
			}

			key.TasksTerminated++
			if key.ProcessId == currentProcess {
				// Update the progress bar is the process that finished coressponds to the current progress bar
				pBar.Increment()
			}

			if key.TasksTerminated >= key.TasksCount {
				// The process has finished, remove it from the map
				delete(processes, task.ProcessId)
				if key.ProcessId != currentProcess {
					// A processor which has no progress bar has finished
					break
				}
				// If the current process has finished, restart a new progress bar
				pBar.Stop()
				currentProcess = ""
				for k := range processes {
					// The first progress bar is the first map entry
					currentProcess = k
					pBar, _ = pterm.DefaultProgressbar.WithTotal(processes[k].TasksCount).WithTitle(processes[k].Name).WithShowElapsedTime(true).Start()
					pBar.Add(processes[k].TasksTerminated)
					break
				}
			}


		}
		if startedProcesses == nil && endedTasks == nil {
			break
		}
	}

	done <- true
}
