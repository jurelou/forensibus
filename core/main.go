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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Ping(ctx, &proto.PingRequest{Identifier: "localhost"})
	if err != nil {
		utils.Log.Fatalf("could not ping worker @%s: %v", serverAddress, err)
	}
	utils.Log.Infof("Worker @%s is up!: %v", serverAddress, r)
}
