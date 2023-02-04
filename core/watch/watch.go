package watch

import (
	"math"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/pterm/pterm"

	dsl "github.com/jurelou/forensibus/core"
	"github.com/jurelou/forensibus/core/run"
	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/names_generator"
)

type onWriteCallback func(fsnotify.Event)

var (
	serverAddress = "localhost:50051"
	mu            sync.Mutex
	waitFor       = 5 * time.Second
	timers        = make(map[string]*time.Timer)
)

func monitorResults(startedProcesses <-chan run.ProcessStarted, endedTasks <-chan run.TaskEnded) {
	for {
		select {
		// case proc, ok := <-startedProcesses:
		// 	if !ok {
		// 		break
		// 	}
		// 	pterm.Success.Println("START", proc)
		case task, ok := <-endedTasks:
			if !ok {
				break
			}
			if utils.IsErrored(task.Status) {
				pterm.Error.Printfln("Processor %s failed (%d errors) against %s: %v", task.Name, len(task.Errors), task.Source, task.Errors)
			} else {
				pterm.Success.Printfln("Successfully ran %s against %s", task.Name, task.Source)
			}
		}
	}
}

func onFileModify(event fsnotify.Event) {
	mu.Lock()
	t, ok := timers[event.Name]
	mu.Unlock()
	if ok {
		t.Reset(waitFor)
	} else {
		pterm.Warning.Printfln("Modified already existing file `%s`, skipped", event.Name)
	}
}

func onFileCreate(watcher *fsnotify.Watcher, event fsnotify.Event, cb onWriteCallback) {
	fileInfo, err := os.Stat(event.Name)
	if err != nil {
		pterm.Warning.Printfln("Error while retrieving new file %s stats: %s", event.Name, err.Error())
		return
	}
	if fileInfo.IsDir() {
		err = watcher.Add(event.Name)
		if err != nil {
			utils.Log.Errorf("Could not watch path %s: %s", event.Name, err.Error())
		}
		pterm.Info.Printfln("Watching new directory %s", event.Name)
		return
	}

	t := time.AfterFunc(math.MaxInt64, func() {
		cb(event)
		if _, ok := timers[event.Name]; ok {
			mu.Lock()
			delete(timers, event.Name)
			mu.Unlock()
		}
	})
	t.Stop()
	mu.Lock()
	timers[event.Name] = t
	mu.Unlock()
	timers[event.Name].Reset(waitFor)
	pterm.Info.Printfln("New file detected: `%s`", event.Name)
}

func watchEvents(watcher *fsnotify.Watcher, cb onWriteCallback) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Has(fsnotify.Create) {
				onFileCreate(watcher, event, cb)
			} else if event.Has(fsnotify.Write) {
				onFileModify(event)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			pterm.Error.Printfln("Error while watching: %s", err.Error())
		}
	}
}

func Watch(pipelineconfigFile string, paths []string, tag string, disableLocalWorker bool) {
	err := utils.ConfigureLogger(disableLocalWorker)
	if err != nil {
		pterm.Error.Printfln(err.Error())
		return
	}

	if tag == "" {
		tag = names_generator.GetRandomName()
	}

	config, err := dsl.LoadAndLint(pipelineconfigFile)
	if err != nil {
		pterm.Error.Printfln(err.Error())
		return
	}

	startedProcesses := make(chan run.ProcessStarted, 512)
	startedTasks := make(chan run.TaskStarted, 512)
	endedTasks := make(chan run.TaskEnded, 512)

	workers, err := run.MakeWorkers(disableLocalWorker, startedTasks, endedTasks)
	if err != nil {
		pterm.Error.Printfln(err.Error())
		return
	}

	if workers.Capacity() == 0 {
		pterm.Error.Println("No workers available, you can either remove -d flag or start a worker")
		return
	} else {
		pterm.Success.Printfln("Launched a pool of %d workers, total capacity is %d", workers.Size(), workers.Capacity())
	}
	pterm.Info.Printfln("Running pipeline `%s`", config.Pipeline.Name)
	pterm.Info.Printfln("Using tag `%s`", tag)
	pterm.Info.Printfln("Using splunk index `%s`", utils.Config.Splunk.Index)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		utils.Log.Errorf("Could not start watcher: %s", err.Error())
	}
	defer watcher.Close()

	go monitorResults(startedProcesses, endedTasks)
	go watchEvents(watcher, func(event fsnotify.Event) {
		fileInfo, err := os.Stat(event.Name)
		if err != nil {
			pterm.Error.Printfln("Error while openning %s", event.Name)
			return
		}
		if fileInfo.Size() == 0 {
			pterm.Warning.Printfln("Skip empty file %s", event.Name)
			return
		}
		inputs, err := run.MakeInputs([]string{event.Name})
		if err != nil {
			pterm.Error.Printfln("Could not create steps from %s: %s", event.Name, err.Error())
			return
		}
		run.RunPipeline(config.Pipeline, inputs, tag, startedProcesses, startedTasks)
	})

	for _, path := range paths {
		err := filepath.Walk(path, func(lpath string, info os.FileInfo, err error) error {
			if err != nil {
				pterm.Warning.Printfln("Error while reading directory %s: %s", lpath, err.Error())
				return err
			}
			if info.IsDir() {
				err = watcher.Add(lpath)
				if err != nil {
					pterm.Error.Printfln("Could not watch path %s: %s", lpath, err.Error())
				} else {
					pterm.Success.Printfln("Watching folder %s", lpath)
				}
			}
			return nil
		})
		if err != nil {
			pterm.Error.Printfln("Could not walk directory %s: %s", path, err.Error())
		}
	}
	// Block main goroutine forever.
	<-make(chan struct{})
}
