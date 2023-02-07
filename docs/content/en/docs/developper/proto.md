---
title: "Protocol buffers"
description: "Forensibus uses GRPC for jobs distribution, which ."
lead: "Forensibus uses GRPC for jobs distribution."
draft: false
images: []
menu:
  docs:
    parent: "Developer guide"
weight: 320
toc: true
---


{{< alert icon="💡" text="Since forensibus uses GRPC, messages are serialized via protobuf." />}}

## Install dependencies

To use protobuf, you need to install the following dependencies:

# protoc >= 3.15

```bash
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v21.12/protoc-21.12-linux-x86_64.zip
unzip protoc-21.12-linux-x86_64.zip -d $HOME/.local
export PATH="$PATH:$HOME/.local/bin"
```

# Go grpc code generator

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

See also the protobuf docs: [protobuf](https://grpc.io/docs/protoc-installation/).

## Update the protocol buffers

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/worker/worker.proto
```

{{< alert icon="💡" text="You can instead run `make proto`." />}}
