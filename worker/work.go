package worker

import (
	"context"
	"fmt"

	"github.com/jurelou/forensibus/proto/worker"
	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

func (*Server) Work(_ context.Context, in *worker.WorkRequest) (*worker.WorkResponse, error) {
	procName := in.GetProcessor()
	source := in.GetSource()
	config := in.GetConfig()
	tag := in.GetTag()
	utils.Log.Infof("Got a task (%s): %s with config %v", procName, source, config)

	// Try to load the processor
	processor, err := processors.Get(procName)
	if err != nil {
		utils.Log.Warnf("Could not find processor %s", procName)
		return &worker.WorkResponse{Status: utils.Failure, Errors: []string{err.Error()}}, err
	}

	// Configure output writer
	// TODO: make writer global
	out := writer.New()
	out.SetTag(tag)
	out.SetDefaultSourceType("forensibus:" + procName)
	out.SetDefaultSource(source)
	out.SetDefaultIndex(utils.Config.Splunk.Index)

	// Run the processor
	pErrors := processor.Run(source, &processors.Config{RawConfig: config}, out)
	if !pErrors.Empty() {
		errStrings := pErrors.AsStrings()
		utils.Log.Warnf("Got %d errors while running %s against %s: %v", pErrors.Len(), procName, source, errStrings)

		formattedErr := fmt.Sprintf("{\"processor\": \"%s\", \"errors\": %s}", procName, pErrors.AsJson())
		ev := writer.NewEvent(formattedErr)
		ev.SourceType = "forensibus_errors"
		out.WriteEvent(ev)

		out.Close()
		return &worker.WorkResponse{Status: utils.Failure, Errors: errStrings}, nil

	}

	out.Close()
	return &worker.WorkResponse{Status: utils.Success, Errors: []string{}}, nil
}
