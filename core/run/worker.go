package run

import (
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	dsl "github.com/jurelou/forensibus/core"
	"github.com/jurelou/forensibus/proto/worker"

	// _ "github.com/jurelou/forensibus/processors/linux"
	// _ "github.com/jurelou/forensibus/processors/linux/commands"
	_ "github.com/jurelou/forensibus/processors/windows"
	_ "github.com/jurelou/forensibus/processors/windows/artifacts"
	_ "github.com/jurelou/forensibus/processors/windows/commands"
)

type Job struct {
	Name   string
	Step   dsl.Step
	Config map[string]string
}

type JobResult struct {
	Status uint32
	Error  string
	Job    Job
}

type Worker struct {
	Client   worker.WorkerClient
	Name     string
	Address  string
	Capacity uint32
}

func (w *Worker) Connect(address string) error {
	w.Address = address
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("Could not connect to %s: %w", address, err)
	}

	w.Client = worker.NewWorkerClient(conn)
	pong, err := w.Ping(3)
	if err != nil {
		return err
	}
	w.Name = pong.GetName()
	w.Capacity = pong.GetCapacity()
	return nil
}

func (w *Worker) Ping(timeout int) (worker.Pong, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	var pong *worker.Pong

	pong, err := w.Client.Ping(ctx, &worker.PingRequest{Identifier: "forensibusCore"}) // TODO: get hostname here
	if err != nil {
		return worker.Pong{}, fmt.Errorf("Could not ping worker %s: %w", w.Address, err)
	}
	return *pong, nil
}

func (w *Worker) Work(wg *sync.WaitGroup, chans JobChannels) {
	defer wg.Done()
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)

	for job := range chans.Jobs {
		fmt.Println("Got job", job.Step.NextArtifact, " FROM ", job.Step.CurrentFolder)

		res, err := w.Client.Work(ctx, &worker.WorkRequest{Source: job.Step.NextArtifact, OutputFolder: job.Step.CurrentFolder, Processor: job.Name, Config: job.Config})
		if err != nil {
			fmt.Println(">>", err)
		}
		chans.JobResults <- JobResult{Job: job, Status: res.GetStatus(), Error: res.GetError()}
	}
	fmt.Println("Worker exited")
}
