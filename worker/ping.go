package worker

import (
	"os"
	"context"
	// "time"
	"github.com/jurelou/forensibus/proto/worker"
	"github.com/jurelou/forensibus/utils"
)

func (s *Server) Ping(ctx context.Context, in *worker.PingRequest) (*worker.Pong, error) {
	utils.Log.Infof("Got a ping from %s", in.GetIdentifier())
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
		utils.Log.Infof("Could not get local hostname %w", err)
	}
	// time.Sleep(10 * time.Second)
	return &worker.Pong{Name: hostname, Capacity: 16}, nil
}
