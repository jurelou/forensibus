package run

import (
	"sync"
	// "fmt"
	// "context"
	// "time"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
	// "github.com/jurelou/forensibus/proto/worker"
)

type WorkerPool struct {
	// PoolSize int
	Workers []Worker
	Wg      sync.WaitGroup
}

func (p *WorkerPool) Connect(address string) (Worker, error) {
	worker := new(Worker)
	err := worker.Connect(address)
	if err != nil {
		return Worker{}, err
	}

	p.Workers = append(p.Workers, *worker)
	return *worker, nil
}

func (p *WorkerPool) Work(chans JobChannels) {
	for _, worker := range p.Workers {
		for i := uint32(0); i < worker.Capacity; i++ {
			p.Wg.Add(1)
			go worker.Work(&p.Wg, chans)
		}
	}
}

func (p *WorkerPool) Size() int {
	return len(p.Workers)
}

func (p *WorkerPool) Capacity() int {
	totalCap := 0
	for _, worker := range p.Workers {
		totalCap += int(worker.Capacity)
	}
	return totalCap
}

func (p *WorkerPool) Wait() {
	p.Wg.Wait()
}
