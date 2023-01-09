package worker

import (
	"context"
	"github.com/jurelou/forensibus/proto/worker"
	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
)

func (s *Server) Work(ctx context.Context, in *worker.WorkRequest) (*worker.WorkResponse, error) {
	procName := in.GetProcessor()
	source := in.GetSource()
	utils.Log.Infof("Got a task (%s): %s", procName, source)

	// Try to load the processor
	processor, err := processors.Get(procName)
	if err != nil {
		return &worker.WorkResponse{Status: utils.Failure, Error: err.Error()}, nil
	}

	// Run the processor
	processor.Run(source)

	return &worker.WorkResponse{Status: utils.Success, Error: ""}, nil
}
