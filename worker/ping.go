package worker

import (
	"context"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/proto"
)

func (s *Server) Ping(ctx context.Context, in *proto.PingRequest) (*proto.PingResponse, error) {
	utils.Log.Infof("Got a ping from %s", in.GetIdentifier())
	// time.Sleep(2 * time.Second)
	return &proto.PingResponse{Message: "Hello " + in.GetIdentifier()}, nil
}
