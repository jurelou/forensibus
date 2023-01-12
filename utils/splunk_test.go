package utils_test

import (
	// "strings"
	"fmt"
	"testing"

	"github.com/jurelou/forensibus/utils"
)

func TestCreateIndex(t *testing.T) {
	// magic := utils.GetMagic("../datasets/archives/Simple.7z")
	// if magic != "application/x-7z-compressed" {
	// 	t.Errorf("Invalid magic for Simple.7z")
	// }

	s := utils.SplunkManagement{DomainName: utils.Config.Splunk.DomainName, Port: utils.Config.Splunk.ManagementPort, Username: utils.Config.Splunk.Username, Password: utils.Config.Splunk.Password}
	err := s.CreateIndex("totqsdaqsdqsdoa")
	fmt.Println("-->", err)

}
