package core

import (
	"context"
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

func walkPipeline(config Config, source []string) {
	utils.Log.Infof("%w", config)
}

func Run(pipelineconfigFile string, paths []string) {
	utils.Log.Info("Starting client")
	pipeline, err := LoadDSLFile(pipelineconfigFile)
	if err != nil {
		utils.Log.Warnf(err.Error())
		return
	}

	SetWorker(serverAddress)

	walkPipeline(pipeline, paths)

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
