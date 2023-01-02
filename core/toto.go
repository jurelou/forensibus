package core

import (
	"fmt"
	"log"
	"time"
	"runtime"
	"github.com/jurelou/forensibus/utils"
)

type T = interface{}

type WorkerPool interface {
	Run()
	AddTask(task func())
}

type workerPool struct {
	maxWorker   int
	queuedTaskC chan func()
}

func NewWorkerPool(maxWorker int) WorkerPool {
	wp := &workerPool{
		maxWorker:   maxWorker,
		queuedTaskC: make(chan func()),
	}

	return wp
}

func (wp *workerPool) Run() {
	wp.run()
}

func (wp *workerPool) AddTask(task func()) {
	wp.queuedTaskC <- task
}

func (wp *workerPool) GetTotalQueuedTask() int {
	return len(wp.queuedTaskC)
}

func (wp *workerPool) run() {
	for i := 0; i < wp.maxWorker; i++ {
		wID := i + 1
		log.Printf("[WorkerPool] Worker %d has been spawned", wID)

		go func(workerID int) {
			for task := range wp.queuedTaskC {
				log.Printf("[WorkerPool] Worker %d start processing task", wID)
				task()
				log.Printf("[WorkerPool] Worker %d finish processing task", wID)
			}
		}(wID)
	}
}

func Yo(pipelineconfigFile string, paths []string) {
	fmt.Println("hello world", paths, "===", pipelineconfigFile)
	config, err := LoadDSLFile(pipelineconfigFile); if err != nil {
		fmt.Println(err,config)
	}
	files, err2 := utils.FindFiles(utils.FindFilesParams{Path: "/root/.profile", PathPatterns: []string{"a", "b"}, FileMagics: []string{".*a", "a", "(.*)a"}})
	if err2 != nil {
		fmt.Println("Error while finding files:", err2, "==", files)
	}

	go func() {
		for {
			time.Sleep(10 * time.Second)
			log.Printf("[main] Total current goroutine: %d", runtime.NumGoroutine())
		}
	}()

	// totalWorker := 5
	// wp := NewWorkerPool(totalWorker)
	// wp.Run()

	// type result struct {
	// 	id    int
	// 	value int
	// }

	// totalTask := 10
	// resultC := make(chan result, totalTask)

	// for i := 0; i < totalTask; i++ {
	// 	id := i + 1
	// 	wp.AddTask(func() {
	// 		log.Printf("[main] Starting task %d", id)
	// 		time.Sleep(5 * time.Second)
	// 		resultC <- result{id, id * 2}
	// 	})
	// }

	// for i := 0; i < totalTask; i++ {
	// 	res := <-resultC
	// 	log.Printf("[main] Task %d has been finished with result %d", res.id, res.value)
	// }

	// waitC := make(chan bool)
	// <-waitC

}