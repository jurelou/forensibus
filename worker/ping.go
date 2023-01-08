package worker

import (
	"context"
	// "time"
	"github.com/jurelou/forensibus/proto"
	"github.com/jurelou/forensibus/utils"
)

func (s *Server) Ping(ctx context.Context, in *proto.PingRequest) (*proto.PingResponse, error) {
	utils.Log.Infof("Got a ping from %s", in.GetIdentifier())
	// time.Sleep(10 * time.Second)
	return &proto.PingResponse{Message: "Hello " + in.GetIdentifier()}, nil
}
