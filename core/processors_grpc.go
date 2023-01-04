package core

import (
	"fmt"
	// "io/ioutil"
	// "log"
	"github.com/hashicorp/go-plugin"
	"github.com/jurelou/forensibus/utils"
	"os"
	"os/exec"
)

func TotoGo() {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: utils.Handshake,
		Plugins:         utils.PluginMap,
		Cmd:             exec.Command("sh", "-c", "../toto"),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("kv_grpc")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	kv := raw.(utils.KV)
	res, err := kv.Get("lol")
	fmt.Println(">>>", res, err)
}
