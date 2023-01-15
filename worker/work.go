package worker

import (
	"context"

	"github.com/jurelou/forensibus/proto/worker"
	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

func (s *Server) Work(ctx context.Context, in *worker.WorkRequest) (*worker.WorkResponse, error) {
	procName := in.GetProcessor()
	source := in.GetSource()
	utils.Log.Infof("Got a task (%s): %s", procName, source)

	// Try to load the processor
	processor, err := processors.Get(procName)
	if err != nil {
		return &worker.WorkResponse{Status: utils.Failure, Error: err.Error()}, err
	}

	// Configure output writer
	// TODO: make writer global
	out := writer.New()
	defer out.Close()
	out.SetDefaultSource(source)
	out.SetDefaultIndex("main")

	// Run the processor
	if err := processor.Run(source, out); err != nil {
		// fmt.Println("oops", err)
		return &worker.WorkResponse{Status: utils.Failure, Error: err.Error()}, nil
	}

	return &worker.WorkResponse{Status: utils.Success, Error: ""}, nil
}
