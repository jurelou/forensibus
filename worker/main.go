package main

import (
	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/zap_logger"
)

func main() {

}

func init() {
	zap_logger.RegisterLogger()
	utils.Log.Info("Salut")
}
