package zap_logger

import (
	"log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/jurelou/forensibus/utils"
)

func RegisterLogger() error {
	zLogger, err := configureLogger()
	if err != nil {
		log.Println("Could not configure logger")
		return err
	}
	defer zLogger.Sync()
	zSugarlog := zLogger.Sugar()
	utils.SetLogger(zSugarlog)
	return nil
}

func configureLogger() (zap.Logger, error) {

	//standard configuration
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.OutputPaths = []string{"forensibus.log", "stdout"}
	// cfg.ErrorOutputPaths = []string{"forensibus.log", "stderr"}

	zLogger, err := cfg.Build()
	if err != nil {
		return *zLogger, err
	}
	return *zLogger, nil
}
