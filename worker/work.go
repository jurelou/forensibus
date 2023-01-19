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
	config := in.GetConfig()
	utils.Log.Infof("Got a task (%s): %s with config %v", procName, source, config)

	// Try to load the processor
	processor, err := processors.Get(procName)
	if err != nil {
		utils.Log.Warnf("Could not find processor %s", procName)
		return &worker.WorkResponse{Status: utils.Failure, Error: err.Error()}, err
	}

	// Configure output writer
	// TODO: make writer global
	out := writer.New()
	// defer out.Close()
	out.SetDefaultSource(source)
	out.SetDefaultIndex(utils.Config.Splunk.Index)

	// Run the processor
	if err := processor.Run(source, out); err != nil {
		utils.Log.Warnf("Error while running %s against %s: %s", procName, source, err.Error())
		out.Close()
		return &worker.WorkResponse{Status: utils.Failure, Error: err.Error()}, nil
	}
	out.Close()
	return &worker.WorkResponse{Status: utils.Success, Error: ""}, nil
}
