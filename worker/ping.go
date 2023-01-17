package worker

import (
	"context"
	"os"

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
		return nil, err
	}
	// time.Sleep(10 * time.Second)
	utils.Log.Infof("qsss %d", utils.Config.WorkersCount)
	return &worker.Pong{Name: hostname, Capacity: utils.Config.WorkersCount}, nil
}
