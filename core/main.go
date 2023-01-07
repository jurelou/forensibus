package core

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/jurelou/forensibus/proto"
	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/decompress"
)

var (
	Client        proto.WorkerClient
	serverAddress = "localhost:50051"
)

// Inbound type is used to track the files being processed (ArtifactPath) and the folder containing files (CurrentFolder)
type Inbound struct {
	ArtifactPath  string
	CurrentFolder string
}

func SetWorker(address string) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.Log.Fatalf("Could not resolve %s: %v", address, err)
	}
	defer conn.Close()

	Client = proto.NewWorkerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err = Client.Ping(ctx, &proto.PingRequest{Identifier: "localhost"})
	if err != nil {
		utils.Log.Fatalf("could not ping worker @%s: %v", address, err)
	}
	utils.Log.Infof("Worker @%s is up!", address)

}

func find(ins []Inbound, config FindConfig) ([]Inbound, error) {
	var outs []Inbound
	var latestErr error = nil

	for _, in := range ins {
		files, err := utils.FindFiles(utils.FindFilesParams{Path: in.ArtifactPath, PathPatterns: config.Patterns, FileFormats: config.MimeTypes})
		latestErr = err
		for _, file := range files {
			outs = append(outs, Inbound{ArtifactPath: file, CurrentFolder: in.CurrentFolder})
		}
	}
	return outs, latestErr
}

func extract(ins []Inbound, config ExtractConfig) ([]Inbound, error) {
	// err :=
	var outs []Inbound
	for _, in := range ins {
		fmt.Println("E>", in)
		out, err := decompress.Decompress(in.ArtifactPath, in.CurrentFolder)
		if err != nil {
			utils.Log.Warnf("Error while decompressing %s: %s", in.ArtifactPath, err.Error())
			continue
		}
		outs = append(outs, Inbound{ArtifactPath: out, CurrentFolder: out})

	}
	return outs, nil
}

func process(ins []Inbound, config ProcessConfig) error {
	utils.Log.Infof("Process %s against %s", config.Name, ins)
	return nil
}

func walk(item interface{}, ins []Inbound) {
	switch t := item.(type) {
	case PipelineConfig:
		// Do nothing with input, simple fanout
		pipelineConfig := item.(PipelineConfig)

		for _, nestedFinds := range pipelineConfig.Finds {
			walk(nestedFinds, ins)
		}
		for _, nestedProcess := range pipelineConfig.Processes {
			walk(nestedProcess, ins)
		}
		for _, nestedExtracts := range pipelineConfig.Extracts {
			walk(nestedExtracts, ins)
		}

	case FindConfig:
		findConfig := item.(FindConfig)
		outs, err := find(ins, findConfig)
		if err != nil {
			utils.Log.Warnf(err.Error())
		}
		outLen := len(outs)
		if outLen == 0 {
			utils.Log.Warnf("Found 0 `%s` files", findConfig.Name)
			break
		} else {
			utils.Log.Infof("Found %d `%s` files", outLen, findConfig.Name)
		}

		for _, nestedFinds := range findConfig.Finds {
			walk(nestedFinds, outs)
		}
		for _, nestedProcess := range findConfig.Processes {
			walk(nestedProcess, outs)
		}
		for _, nestedExtracts := range findConfig.Extracts {
			walk(nestedExtracts, outs)
		}
	case ExtractConfig:
		extractConfig := item.(ExtractConfig)
		out, err := extract(ins, extractConfig)
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
			walk(nestedFinds, out)
		}
		for _, nestedProcess := range extractConfig.Processes {
			walk(nestedProcess, out)
		}
		for _, nestedExtracts := range extractConfig.Extracts {
			walk(nestedExtracts, out)
		}
	case ProcessConfig:
		processConfig := item.(ProcessConfig)
		err := process(ins, processConfig)
		if err != nil {
			utils.Log.Warnf("Error while processing: %s", err.Error())
		}
	default:
		fmt.Printf("I don't know about type %T!\n", t)
	}

}

func walkPipeline(config Config, sources []string) {
	pipeline := config.Pipeline
	utils.Log.Infof("Running pipeline `%s`", pipeline.Name)

	ins := make([]Inbound, len(sources))
	for _, source := range sources {
		absPath, _ := utils.ConvertRelativePath(source)
		filename := filepath.Base(absPath)

		filenameWhithoutSuffix := strings.TrimSuffix(filename, filepath.Ext(filename))
		outputFolder := filepath.Join(utils.Config.OutputFolder, filepath.Dir(absPath), filenameWhithoutSuffix)

		utils.Log.Infof("Writing output filed to %s", outputFolder)
		ins = append(ins, Inbound{ArtifactPath: absPath, CurrentFolder: outputFolder})
	}
	walk(pipeline, ins)

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
