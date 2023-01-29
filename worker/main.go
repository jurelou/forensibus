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

func RunWorkerServer(srv *grpc.Server, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %w", port, err)
	}

	loadProcessors()

	s := grpc.NewServer()
	worker.RegisterWorkerServer(s, &Server{})
	utils.Log.Infof("Worker listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}
	return nil
}

func Run(workersCount uint32) {
	err := utils.ConfigureLogger(true)
	if err != nil {
		utils.Log.Errorf(err.Error())
		return
	}

	viper.Set("WorkersCount", workersCount)
	utils.Config.WorkersCount = workersCount
	utils.Log.Infof("Launched with %d workers", utils.Config.WorkersCount)

	var port = 50051
	var server *grpc.Server

	err = RunWorkerServer(server, port)
	if err != nil {
		utils.Log.Errorf("Could not start worker server: %s", err.Error())
	}
}
