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
	sourcetype := in.GetSourcetype()
	rawConfig := in.GetConfig()
	tag := in.GetTag()
	utils.Log.Infof("Got a task (%s): %s with config %v", procName, source, rawConfig)

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
	out.SetDefaultSource(source)
	out.SetDefaultIndex(in.GetSplunkIndex())

	if sourcetype != "" {
		out.SetDefaultSourceType(sourcetype)
	} else {
		out.SetDefaultSourceType("forensibus:" + procName)
	}

	// Run the processor
	pErrors := processor.Run(source, &processors.Config{RawConfig: rawConfig}, out)
	if !pErrors.Empty() {
		errStrings := pErrors.AsStrings()
		utils.Log.Warnf("Got %d errors while running %s against %s: %v", pErrors.Len(), procName, source, errStrings)

		formattedErr := fmt.Sprintf("{\"processor\": \"%s\", \"errors\": %s}", procName, pErrors.AsJson())
		ev := writer.NewEvent(formattedErr)
		ev.SourceType = "forensibus:processor_errors"
		out.WriteEvent(ev)

		out.Close()
		return &worker.WorkResponse{Status: utils.Failure, Errors: errStrings}, nil

	}

	out.Close()
	return &worker.WorkResponse{Status: utils.Success, Errors: []string{}}, nil
}
