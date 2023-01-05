package worker

import (
	_ "context"
	"fmt"
	"net"
	_ "time"

	"github.com/jurelou/forensibus/proto"
	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"

	"google.golang.org/grpc"
)

var (
	port = 50051
)

// server is used to implement proto.WorkerServer.
type Server struct {
	proto.UnimplementedWorkerServer
}

func loadProcessors() {
	for procName, proc := range processors.Registry {
		proc.Configure()
		utils.Log.Infof("Configured processor %s", procName)
	}
}

func runGRPCServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		utils.Log.Fatalf("failed to listen on port %d: %v", port, err)
	}

	loadProcessors()

	s := grpc.NewServer()
	proto.RegisterWorkerServer(s, &Server{})
	utils.Log.Infof("Worker listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		utils.Log.Fatalf("failed to serve: %v", err)
	}
}

func Run() {
	runGRPCServer()
}
