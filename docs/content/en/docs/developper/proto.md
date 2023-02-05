---
title: "Protocol buffers"
description: "Forensibus uses GRPC for jobs distribution, which ."
lead: "Forensibus uses GRPC for jobs distribution."
draft: false
images: []
menu:
  docs:
    parent: "Developer guide"
weight: 130
toc: true
---


{{< alert icon="💡" text="Since forensibus uses GRPC, messages are serialized via protobuf." />}}

## Install dependencies

To use protobuf, you need to install the following dependencies:

```bash
sudo apt install -y protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

See also the protobuf docs: [protobuf](https://grpc.io/docs/protoc-installation/).

## Update the protocol buffers


```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/worker/worker.proto
```

{{< alert icon="💡" text="You can instead run `make proto`." />}}
