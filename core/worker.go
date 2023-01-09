package core

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/jurelou/forensibus/proto/worker"

	// _ "github.com/jurelou/forensibus/processors/linux"
	// _ "github.com/jurelou/forensibus/processors/linux/commands"
	_ "github.com/jurelou/forensibus/utils/processors"

	_ "github.com/jurelou/forensibus/processors/windows"
	_ "github.com/jurelou/forensibus/processors/windows/artifacts"
	_ "github.com/jurelou/forensibus/processors/windows/commands"
)

var (
	queueSize     = 4
	jobsChannel   = make(chan Job, queueSize)
	resultChannel = make(chan JobResult, queueSize)
)

type Worker struct {
	// PoolSize int
	Client  worker.WorkerClient
	Name    string
	Address string
}

func (w *Worker) Connect(address string) error {
	w.Address = address
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("Could not connect to %s: %w", address, err)
	}
	w.Client = worker.NewWorkerClient(conn)
	pingResponse, err := w.Ping(3)
	if err != nil {
		return err
	}
	w.Name = pingResponse.GetMessage()
	return nil
}

func (w *Worker) Ping(timeout int) (worker.PingResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	var response *worker.PingResponse

	response, err := w.Client.Ping(ctx, &worker.PingRequest{Identifier: "forensibusCore"}) // TODO: get hostname here
	if err != nil {
		return *response, fmt.Errorf("Could not ping worker %s: %w", w.Address, err)
	}
	return *response, nil
}

func (w *Worker) Work(jobs <-chan Job, results chan<- JobResult) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)

	for job := range jobs {
		fmt.Println("Got job", job.In.Next)

		res, err := w.Client.Work(ctx, &worker.WorkRequest{Source: job.In.Next, OutputFolder: job.In.Current, Processor: job.Config.Name, Config: job.Config.Config})
		if err != nil {
			fmt.Println(">>", err)
		}
		results <- JobResult{Processor: job.Config.Name, In: job.In, Status: res.GetStatus(), Error: res.GetError()}
	}

}

type Job struct {
	In     Input
	Config ProcessConfig
	// Results chan JobResult
}

type JobResult struct {
	Status    uint32
	Error     string
	In        Input
	Processor string
}

// func work(jobsChannel <-chan Job) {
// 	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)

// 	for job := range jobsChannel {
// 		fmt.Println("Got job", job.In.Next)

// 		res, err := Client.Work(ctx, &worker.WorkRequest{Source: job.In.Next, OutputFolder: job.In.Current, Processor: job.Config.Name, Config: job.Config.Config})
// 		if err != nil {
// 			fmt.Println(">>", err)
// 		}
// 		job.Results <- JobResult{Processor: job.Config.Name, In: job.In, Status: res.GetStatus(), Error: res.GetError()}
// 	}
// }

func StartWorkers() {
	// for w := 0; w < queueSize; w++ {
	// 	go work(jobsChannel)
	// }
}

func RunProcessor(ins []Input, config ProcessConfig) {
	// results := make(chan string, len(ins))
	// resultsChan := make(chan JobResult, len(ins))

	// job := Job{ins: ins, config: config, results: resultsChan}
	// jobsChannel <- job

	fmt.Printf("Launching %d jobs\n", len(ins))
	for _, in := range ins {
		fmt.Println("Launch ", in.Next)
		jobsChannel <- Job{In: in, Config: config}
	}

	for i := 0; i < len(ins); i++ {
		res := <-resultChannel
		fmt.Println("Response:", res.Processor, res.Status, res.Error)
	}

}
