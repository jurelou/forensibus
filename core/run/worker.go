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
	"github.com/jurelou/forensibus/utils"

	// _ "github.com/jurelou/forensibus/processors/linux"
	// _ "github.com/jurelou/forensibus/processors/linux/commands"
	_ "github.com/jurelou/forensibus/processors"
	// _ "github.com/jurelou/forensibus/processors/windows"
	_ "github.com/jurelou/forensibus/processors/windows/ese"
	_ "github.com/jurelou/forensibus/processors/windows/artifacts"
	_ "github.com/jurelou/forensibus/processors/windows/commands"
)

type Job struct {
	Name      string
	Tag       string
	ProcessId string
	Step      dsl.Step
	Config    map[string]string
}

type JobResult struct {
	Status uint32
	Errors []string
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
		return fmt.Errorf("could not connect to %s: %w", address, err)
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

func (w *Worker) Ping(timeout int) (*worker.Pong, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	for attempts := 0; attempts < 4; attempts++ {
		pong, err := w.Client.Ping(ctx, &worker.PingRequest{Identifier: "forensibusCore"}) // TODO: get hostname here
		if err == nil && pong != nil {
			return pong, nil
		}
		time.Sleep(1 * time.Second)
	}
	return nil, fmt.Errorf("could not ping worker at %s", w.Address)
}

func (w *Worker) Work(wg *sync.WaitGroup, tStart <-chan TaskStarted, tEnd chan<- TaskEnded) {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(utils.Config.ProcessorTimeout)*time.Second)
	defer cancel()
	for task := range tStart {
		res, err := w.Client.Work(ctx, &worker.WorkRequest{
			Source:       task.Step.NextArtifact,
			OutputFolder: task.Step.CurrentFolder,
			Processor:    task.ProcessConfig.Name,
			Config:       task.ProcessConfig.Config,
			Tag:          task.Tag,
		})
		jobErrors := res.GetErrors()
		jobStatus := res.GetStatus()
		if err != nil {
			jobErrors = []string{fmt.Sprintf("Processor %s timed out (%s): %s", task.ProcessConfig.Name, task.Step.NextArtifact, err.Error())}
			jobStatus = utils.Timeout
		}
		tEnd <- TaskEnded{
			Name:      task.ProcessConfig.Name,
			Source:    task.Step.NextArtifact,
			ProcessId: task.ProcessId,
			Status:    jobStatus,
			Errors:    jobErrors,
		}
	}
}
