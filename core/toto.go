package core

import (
	"fmt"
	"log"
	"time"
	"runtime"
	_"github.com/jurelou/forensibus/utils"
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

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func Yo(pipelineconfigFile string, paths []string) {
	fmt.Println("hello world", paths, "===", pipelineconfigFile)
	config, err := LoadDSLFile(pipelineconfigFile); if err != nil {
		fmt.Println(err,config)
	}
	// files, err2 := utils.FindFiles(utils.FindFilesParams{Path: "README.md", PathPatterns: []string{"a", "b", "^[Rr]"}, FileMagics: []string{".*a", "a", "(.*)a"}})
	// if err2 != nil {
	// 	fmt.Println("Error while finding files:", err2, "==", files)
	// }
	// fmt.Println("found files ->", files)

	go func() {
		for {
			time.Sleep(10 * time.Second)
			log.Printf("[main] Total current goroutine: %d", runtime.NumGoroutine())
		}
	}()

	const numJobs = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, 10)


    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

	for j := 1; j <= 10 ; j++ {
		fmt.Println("Push job", j)
        jobs <- j
    }
	// close(jobs)

	for a := 1; a <= 10; a++ {
        <-results
    }

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