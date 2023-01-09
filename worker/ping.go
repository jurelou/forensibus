package worker

import (
	"context"
	// "time"
	"github.com/jurelou/forensibus/proto/worker"
	"github.com/jurelou/forensibus/utils"
)

func (s *Server) Ping(ctx context.Context, in *worker.PingRequest) (*worker.PingResponse, error) {
	utils.Log.Infof("Got a ping from %s", in.GetIdentifier())
	// time.Sleep(10 * time.Second)
	return &worker.PingResponse{Message: "Hello " + in.GetIdentifier()}, nil
}
