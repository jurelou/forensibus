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
	in      Inbound
	config  ProcessConfig
	results chan JobResult
}

type JobResult struct {
	data string
}

func worker(jobsChannel <-chan Job) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)

	for job := range jobsChannel {
		fmt.Println("Got job", job.in.ArtifactPath)
		// time.Sleep(time.Second * 5)
		_, err := Client.Ping(ctx, &proto.PingRequest{Identifier: job.in.ArtifactPath})
		if err != nil {
			fmt.Println(">>", err)
		}
		job.results <- JobResult{data: fmt.Sprintf("done %s", job.in.ArtifactPath)}
	}
}

func StartWorkers() {
	for w := 0; w < queueSize; w++ {
		go worker(jobsChannel)
	}
}

func RunProcessor(ins []Inbound, config ProcessConfig) {
	// results := make(chan string, len(ins))
	resultsChan := make(chan JobResult, len(ins))

	// job := Job{ins: ins, config: config, results: resultsChan}
	// jobsChannel <- job

	fmt.Printf("Launching %d jobs\n", len(ins))
	for _, in := range ins {
		fmt.Println("Launch ", in.ArtifactPath)
		jobsChannel <- Job{in: in, config: config, results: resultsChan}
	}

	for i := 0; i < len(ins); i++ {
		a := <-resultsChan
		fmt.Println("RES:", a.data)
	}

}
