package main

import (
	"fmt"
	"github.com/hashicorp/go-plugin"
	"github.com/jurelou/forensibus/utils"
	"io/ioutil"
)

// Here is a real implementation of KV that writes to a local file with
// the key name and the contents are the value of the key.
type KV struct{}

func (KV) Put(key string, value []byte) error {
	fmt.Println("PUUT")
	value = []byte(fmt.Sprintf("%s\n\nWritten from plugin-go-grpc", string(value)))
	return ioutil.WriteFile("kv_"+key, value, 0644)
}

func (KV) Get(key string) ([]byte, error) {
	fmt.Println("GEET")
	return ioutil.ReadFile("kv_" + key)
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: utils.Handshake,
		Plugins: map[string]plugin.Plugin{
			"kv": &utils.KVGRPCPlugin{Impl: &KV{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
	fmt.Println("INIT MAIN")
}
