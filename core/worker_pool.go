package core

import (
// "context"
// "time"

// "google.golang.org/grpc"
// "google.golang.org/grpc/credentials/insecure"

// "github.com/jurelou/forensibus/proto/worker"
)

type WorkerPool struct {
	// PoolSize int
	Workers []Worker
}

func (p *WorkerPool) Connect(address string) (Worker, error) {
	worker := new(Worker)
	if err := worker.Connect(address); err != nil {
		return *worker, err
	}
	p.Workers = append(p.Workers, *worker)
	return *worker, nil
}

func (p *WorkerPool) Work(jobs <-chan Job, results chan<- JobResult) {
	for _, worker := range p.Workers {
		for i := uint32(0) ; i < worker.Capacity ; i++ {
			go worker.Work(jobs, results)
		}
	}
}

func (p *WorkerPool) Size() int {
	return len(p.Workers)
}

func (p *WorkerPool) Capacity() int {
	cap := 0
	for _, worker := range p.Workers {
		cap += int(worker.Capacity)
	}
	return cap
}
