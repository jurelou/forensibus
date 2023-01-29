package zap_logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetLogger(enableStdout bool) (*zap.SugaredLogger, error) {
	// var sugarLogger *zap.SugaredLogger

	// standard configuration
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	cfg.Encoding = "console"
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if enableStdout {
		cfg.OutputPaths = []string{"forensibus.log", "stdout"}
	} else {
		cfg.OutputPaths = []string{"forensibus.log"}
	}
	// cfg.ErrorOutputPaths = []string{"forensibus.log", "stderr"}

	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("Could not configure zap logger: %w", err)
	}
	defer logger.Sync()

	sugarLogger := logger.Sugar()
	return sugarLogger, nil
}
