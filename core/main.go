package core

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/jurelou/forensibus/proto"
	"github.com/jurelou/forensibus/utils"
)

var (
	Client        proto.WorkerClient
	serverAddress = "localhost:50051"
)

func SetWorker(address string) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.Log.Fatalf("Could not resolve %s: %v", address, err)
	}
	defer conn.Close()

	Client = proto.NewWorkerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	res, err := Client.Ping(ctx, &proto.PingRequest{Identifier: "localhost"})
	if err != nil {
		utils.Log.Fatalf("could not ping worker @%s: %v", address, err)
	}
	utils.Log.Infof("Worker @%s is up! (%s) ", address, res)

}


func find(in []string, config FindConfig) ([]string, error) {
	var out []string
	var latestErr error = nil

	for _, file := range in {
		files, err := utils.FindFiles(utils.FindFilesParams{Path: file, PathPatterns: config.Patterns, FileFormats: config.MimeTypes})
		latestErr = err
		out = append(out, files...)
	}
	return out, latestErr
}

func extract(in []string, config ExtractConfig) ([]string, error) {
	return in, nil
}

func process(in []string, config ProcessConfig) error {

	utils.Log.Infof("Process %s against %s", config.Name, in)
	return nil
}
func walkPipeline(config Config, sources []string) {
	pipeline := config.Pipeline
	utils.Log.Infof("Running pipeline `%s`", pipeline.Name)

	var walkConfig func(item interface{}, in []string)

	walkConfig = func(item interface{}, in []string) {
		switch t := item.(type) {
		case PipelineConfig:
			// Do nothing with input, simple fanout
			pipelineConfig := item.(PipelineConfig)

			for _, nestedFinds := range pipelineConfig.Finds {
				walkConfig(nestedFinds, in)
			}
			for _, nestedProcess := range pipelineConfig.Processes {
				walkConfig(nestedProcess, in)
			}
			for _, nestedExtracts := range pipelineConfig.Extracts {
				walkConfig(nestedExtracts, in)
			}

		case FindConfig:
			findConfig := item.(FindConfig)
			out, err := find(in, findConfig)
			if err != nil {
				utils.Log.Warnf(err.Error())
			}
			outLen := len(out)
			if outLen == 0 {
				utils.Log.Warnf("Found 0 `%s` files", findConfig.Name)
				break
			} else {
				utils.Log.Infof("Found %d `%s` files", outLen, findConfig.Name)
			}

			for _, nestedFinds := range findConfig.Finds {
				walkConfig(nestedFinds, out)
			}
			for _, nestedProcess := range findConfig.Processes {
				walkConfig(nestedProcess, out)
			}
			for _, nestedExtracts := range findConfig.Extracts {
				walkConfig(nestedExtracts, out)
			}
		case ExtractConfig:
			extractConfig := item.(ExtractConfig)
			out, err := extract(in, extractConfig)
			if err != nil {
				utils.Log.Warnf(err.Error())
			}
			outLen := len(out)
			if outLen == 0 {
				utils.Log.Warnf("Extracted 0 `%s` files", extractConfig.Name)
				break
			} else {
				utils.Log.Infof("Extracted %d `%s` files", outLen, extractConfig.Name)
			}
			for _, nestedFinds := range extractConfig.Finds {
				walkConfig(nestedFinds, out)
			}
			for _, nestedProcess := range extractConfig.Processes {
				walkConfig(nestedProcess, out)
			}
			for _, nestedExtracts := range extractConfig.Extracts {
				walkConfig(nestedExtracts, out)
			}
		case ProcessConfig:
			processConfig := item.(ProcessConfig)
			err := process(in, processConfig)
			if err != nil {
				utils.Log.Warnf("Error while processing: %s", err.Error())
			}
		default:
			fmt.Printf("I don't know about type %T!\n", t)
		}

	}

	walkConfig(pipeline, sources)

}

func Run(pipelineconfigFile string, paths []string) {
	utils.Log.Info("Starting client")

	config, err := LoadDSLFile(pipelineconfigFile)
	if err != nil {
		utils.Log.Warnf(err.Error())
		return
	}

	SetWorker(serverAddress)

	walkPipeline(config, paths)

	// r2, err2 := c.Work(ctx, &proto.WorkRequest{Source: "src", Processor: "evtx", Config: map[string]string{}})
	// if err2 != nil {
	// 	utils.Log.Fatalf("could not ping worker @%s: %v", serverAddress, err2)
	// }
	// status := r2.GetStatus()
	// if (status > utils.Success) {
	// 	error := r2.GetError()
	// 	utils.Log.Warnf("Error from processor %s (%d): %s", "aa", status, error)

	// }
}
