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
	serverAddress = "localhost:50051"
)

func Run(pipelineconfigFile string, paths []string) {
	utils.Log.Info("Starting client")

	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.Log.Fatalf("Could not create GRPC client: %v", err)
	}
	defer conn.Close()

	c := proto.NewWorkerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 60 * time.Second)
	defer cancel()

	r, err := c.Ping(ctx, &proto.PingRequest{Identifier: "localhost"})
	if err != nil {
		utils.Log.Fatalf("could not ping worker @%s: %v", serverAddress, err)
	}
	utils.Log.Infof("Worker @%s is up!: %v", serverAddress, r)


	r2, err2 := c.Work(ctx, &proto.WorkRequest{Source: "src", Processor: "evtx", Config: map[string]string{}})
	if err2 != nil {
		utils.Log.Fatalf("could not ping worker @%s: %v", serverAddress, err2)
	}
	status := r2.GetStatus()
	if (status > utils.Success) {
		error := r2.GetError()
		utils.Log.Warnf("Error from processor %s (%d): %s", "aa", status, error)

	}
}
