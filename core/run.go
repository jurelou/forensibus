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

func ConnectWorker(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.Log.Fatalf("Could not resolve %s: %v", address, err)
	}

	Client = proto.NewWorkerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err = Client.Ping(ctx, &proto.PingRequest{Identifier: "localhost"})
	if err != nil {
		return nil, fmt.Errorf("Could not ping worker @%s: %w", address, err)
	}
	utils.Log.Infof("Worker @%s is up!", address)
	return conn, nil
}

func find(ins []Input, patterns []string, mimes []string) ([]Input, error) {
	var outs []Input
	var latestErr error = nil

	for _, in := range ins {
		files, err := utils.FindFiles(utils.FindFilesParams{Path: in.Next, PathPatterns: patterns, FileFormats: mimes})
		latestErr = err
		for _, file := range files {
			outs = append(outs, Input{Next: file, Current: in.Current})
		}
	}
	return outs, latestErr
}

// Finds and extract files matching patterns
func extract(ins []Input, config ExtractConfig) ([]Input, error) {
	// err :=
	var outs []Input
	files, _ := find(ins, config.Patterns, config.MimeTypes)

	for _, in := range files {
		out, err := decompress.Decompress(in.Next, in.Current)
		if err != nil {
			utils.Log.Warnf("Error while decompressing %s: %s", in.Next, err.Error())
			continue
		}
		outs = append(outs, Input{Next: out, Current: out})
	}
	return outs, nil
}

func process(ins []Input, config ProcessConfig) error {
	RunProcessor(ins, config)
	return nil
}

func execPipeline(config Config, in []Input) {

	WalkPipeline(config.Pipeline, in, func(item interface{}, in []Input) ([]Input, error) {
		var err error
		out := []Input{}

		switch item.(type) {
		case FindConfig:
			findConfig := item.(FindConfig)
			out, err = find(in, findConfig.Patterns, findConfig.MimeTypes)
			if err != nil {
				utils.Log.Warnf(err.Error())
				return []Input{}, err
			}
			outLen := len(out)
			if outLen == 0 {
				utils.Log.Warnf("Found 0 `%s` files", findConfig.Name)
			} else {
				utils.Log.Infof("Found %d `%s` files", outLen, findConfig.Name)
			}
			return out, nil

		case ExtractConfig:
			extractConfig := item.(ExtractConfig)
			out, err = extract(in, extractConfig)
			if err != nil {
				utils.Log.Warnf(err.Error())
				return []Input{}, err
			}
			outLen := len(out)
			if outLen == 0 {
				utils.Log.Warnf("Extracted 0 `%s` files", extractConfig.Name)
			} else {
				utils.Log.Infof("Extracted %d `%s` files", outLen, extractConfig.Name)
			}
			return out, nil

		case ProcessConfig:
			processConfig := item.(ProcessConfig)
			err := process(in, processConfig)
			if err != nil {
				utils.Log.Warnf("Error while processing %s: %s", processConfig.Name, err.Error())
			}
			return []Input{}, err

		}
		return in, nil
	})
}

func makeInputs(sources []string) ([]Input, error) {
	ins := make([]Input, 0, len(sources))
	for _, source := range sources {
		absPath, err := utils.ConvertRelativePath(source)
		if err != nil {
			return ins, err
		}

		filename := filepath.Base(absPath)
		filenameWhithoutSuffix := strings.TrimSuffix(filename, filepath.Ext(filename))

		volumeName := filepath.VolumeName(absPath)
		absPath = strings.TrimPrefix(absPath, volumeName)

		outputFolder := filepath.Join(utils.Config.OutputFolder, filepath.Dir(absPath), filenameWhithoutSuffix)

		utils.Log.Infof("Writing output filed FROM `%s` TO `%s`", outputFolder, absPath)
		ins = append(ins, Input{Current: outputFolder, Next: absPath})
	}
	return ins, nil
}

func Run(pipelineconfigFile string, paths []string) {
	utils.Log.Info("Starting client")

	config, err := LoadDSLFile(pipelineconfigFile)
	if err != nil {
		utils.Log.Warnf(err.Error())
		return
	}

	conn, err := ConnectWorker(serverAddress)
	if err != nil {
		utils.Log.Warnf(err.Error())
		return
	}
	defer conn.Close()

	StartWorkers()

	inputs, err := makeInputs(paths)
	if err != nil {
		utils.Log.Warnf(err.Error())
		return
	}
	execPipeline(config, inputs)

}
