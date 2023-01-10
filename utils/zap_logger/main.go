package zap_logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func GetLogger() (*zap.SugaredLogger, error) {
	var sugarLogger *zap.SugaredLogger

	logger, err := configureLogger()
	if err != nil {
		log.Println("Could not configure logger")
		return sugarLogger, err
	}
	defer logger.Sync()

	sugarLogger = logger.Sugar()
	return sugarLogger, nil
}

func configureLogger() (zap.Logger, error) {
	//standard configuration
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	cfg.Encoding = "console"
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.OutputPaths = []string{"forensibus.log", "stdout"}
	// cfg.ErrorOutputPaths = []string{"forensibus.log", "stderr"}

	zLogger, err := cfg.Build()
	if err != nil {
		return zap.Logger{}, err
	}
	return *zLogger, nil
}
