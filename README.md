# forensibus

# TODO
- wrap errors properly
- grpc-plugin
- utils/filesystem: add some file types to the filetype lib. (evtx, registry, text ...)
- cli: convert relative filepath to absolute filepath
- Add a --force flag to the run command, force exec even if some errors appear (some workers are unavailable, ....)
- Find (hcl) Add a `recurse` boolean
- Find (hcl) Add a `type` option file/dir (file by default)
- set windows PE ressources : https://github.com/tc-hib/go-winres
- cli: Add worker ping, plugin list, plugin info, ....

- cli manage flags with viper  + cobra / https://github.com/desdic/playground/blob/master/go/cobraviper/cmd/root.go
- use libmagic bindings instead of filetype 
- Add more archives extraction types
- Ability to lint the config file. (check if processors exists, ....)

# Compiler pour les différents OS

sudo apt-get install mingw-w64-x86-64-dev gcc-mingw-w64-x86-64 gcc-mingw-w64 gcc-multilib

# Mettre à jour le protocol (protobuf)

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/worker.proto
```

# grpc

https://github.com/pahanini/go-grpc-bidirectional-streaming-example
https://github.com/toransahu/grpc-eg-go

# HCL

https://github.com/floresj/hcl2-examples
https://github.com/ivancorrales/hcl-by-example

# workers pool

https://github.com/panjf2000/ants

# Queues

- FIFO: https://gist.github.com/moraes/2141121
- https://github.com/enriquebris/goconcurrentqueue
- https://github.com/hibiken/asynq

# Rabbitmq 

- Direct reply to: https://gist.github.com/Shazambom/66111731ef50df83e7126e6c28aa7c04

# splunk

https://github.com/ZachtimusPrime/Go-Splunk-HTTP

# Logging

https://github.com/uber-go/zap
https://github.com/rs/zerolog

# cli

- progress bar https://github.com/gosuri/uiprogress
- cli: https://github.com/spf13/cobra
- ui: https://github.com/charmbracelet/bubbletea

# Others

- bloom filters https://github.com/bits-and-blooms/bloom
- makefile alternative: https://github.com/magefile/mage
- plugin: https://github.com/hashicorp/go-plugin
- config: https://github.com/spf13/viper
- filesystem: https://github.com/spf13/afero
- dag https://github.com/mostafa-asg/dag
- singletoons: https://github.com/uber-go/fx