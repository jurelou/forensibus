package worker

import (
	"fmt"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/jurelou/forensibus/proto/worker"
	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"

	_ "context"
	_ "time"
)

var port = 50051

// server is used to implement proto.WorkerServer.
type Server struct {
	worker.UnimplementedWorkerServer
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
	worker.RegisterWorkerServer(s, &Server{})
	utils.Log.Infof("Worker listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		utils.Log.Fatalf("failed to serve: %v", err)
	}
}

func Run(workersCount uint32) {
	viper.Set("WorkersCount", workersCount)
	utils.Config.WorkersCount = workersCount
	utils.Log.Infof("Launched with %d workers", utils.Config.WorkersCount)

	runGRPCServer()
}
