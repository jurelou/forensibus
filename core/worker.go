package core

import (
	"context"
	"fmt"
	"time"

	"github.com/jurelou/forensibus/proto"

	// _ "github.com/jurelou/forensibus/processors/linux"
	// _ "github.com/jurelou/forensibus/processors/linux/commands"
	_ "github.com/jurelou/forensibus/utils/processors"

	_ "github.com/jurelou/forensibus/processors/windows"
	_ "github.com/jurelou/forensibus/processors/windows/artifacts"
	_ "github.com/jurelou/forensibus/processors/windows/commands"
)

var (
	queueSize   = 4
	jobsChannel = make(chan Job, queueSize)
)

type Job struct {
	In      Input
	Config  ProcessConfig
	Results chan JobResult
}

type JobResult struct {
	Status    uint32
	Error     string
	In        Input
	Processor string
}

func work(jobsChannel <-chan Job) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)

	for job := range jobsChannel {
		fmt.Println("Got job", job.In.Next)

		res, err := Client.Work(ctx, &proto.WorkRequest{Source: job.In.Next, OutputFolder: job.In.Current, Processor: job.Config.Name, Config: job.Config.Config})
		if err != nil {
			fmt.Println(">>", err)
		}
		job.Results <- JobResult{Processor: job.Config.Name, In: job.In, Status: res.GetStatus(), Error: res.GetError()}
	}
}

func StartWorkers() {
	for w := 0; w < queueSize; w++ {
		go work(jobsChannel)
	}
}

func RunProcessor(ins []Input, config ProcessConfig) {
	// results := make(chan string, len(ins))
	resultsChan := make(chan JobResult, len(ins))

	// job := Job{ins: ins, config: config, results: resultsChan}
	// jobsChannel <- job

	fmt.Printf("Launching %d jobs\n", len(ins))
	for _, in := range ins {
		fmt.Println("Launch ", in.Next)
		jobsChannel <- Job{In: in, Config: config, Results: resultsChan}
	}

	for i := 0; i < len(ins); i++ {
		res := <-resultsChan
		fmt.Println("Response:", res.Processor, res.Status, res.Error)
	}

}
