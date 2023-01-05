package worker

import (
	"context"
	"fmt"
	"net"
	_ "time"

	"github.com/jurelou/forensibus/proto"
	"github.com/jurelou/forensibus/utils"
	"google.golang.org/grpc"
)

var (
	port = 50051
)

// server is used to implement proto.WorkerServer.
type server struct {
	proto.UnimplementedWorkerServer
}

// Ping implements proto.WorkerServer
func (s *server) Ping(ctx context.Context, in *proto.PingRequest) (*proto.PingResponse, error) {
	utils.Log.Infof("Got a ping from %s", in.GetIdentifier())
	// time.Sleep(2 * time.Second)
	return &proto.PingResponse{Message: "Hello " + in.GetIdentifier()}, nil
}

func Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		utils.Log.Fatalf("failed to listen on port %d: %v", port, err)
	}

	s := grpc.NewServer()
	proto.RegisterWorkerServer(s, &server{})
	utils.Log.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		utils.Log.Fatalf("failed to serve: %v", err)
	}
}
