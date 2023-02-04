package watch

import (
	"math"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/pterm/pterm"

	dsl "github.com/jurelou/forensibus/core"
	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/names_generator"
)

func parseFile(event fsnotify.Event) {
	pterm.Info.Printfln("Processing file: `%s`", event.Name)
}

func watchEvents(watcher *fsnotify.Watcher) {
	var (
		mu      sync.Mutex
		waitFor = 5 * time.Second
		timers  = make(map[string]*time.Timer)
	)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			// log.Println("event:", event.Name, event)
			if event.Has(fsnotify.Create) {
				t := time.AfterFunc(math.MaxInt64, func() {
					parseFile(event)
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
			} else if event.Has(fsnotify.Write) {
				mu.Lock()
				t, ok := timers[event.Name]
				mu.Unlock()

				if ok {
					t.Reset(waitFor)
				} else {
					pterm.Warning.Printfln("Modified already existing file `%s`, skipped", event.Name)
				}
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

	pterm.Info.Printfln("Running pipeline `%s` against %v", config.Pipeline.Name, paths)
	pterm.Info.Printfln("Using tag `%s`", tag)
	pterm.Info.Printfln("Using splunk index `%s`", utils.Config.Splunk.Index)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		utils.Log.Errorf("Could not start watcher: %s", err.Error())
	}
	defer watcher.Close()

	go watchEvents(watcher)

	for _, path := range paths {
		err = watcher.Add(path)
		if err != nil {
			utils.Log.Errorf("Could not watch path %s: %s", path, err.Error())
		}
	}
	// Block main goroutine forever.
	<-make(chan struct{})
}
