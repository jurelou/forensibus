package core

import (
	"fmt"
	"testing"
	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/zap_logger"
)

func TestXxssx(t* testing.T) {
	// TotoAa()
	fmt.Println("!!!")
	zap_logger.RegisterLogger()

	utils.Log.Info("salut")
	utils.Log.Debug("salut")
	utils.Log.Errorf("salut")
	// utils.Log.Warnf("salut")

}