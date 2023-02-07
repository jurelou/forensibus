---
title: "Code quality"
description: ""
lead: ""
draft: false
images: []
menu:
  docs:
    parent: "Developer guide"
weight: 320
toc: true
---

# Install dependencies

```bash
go install mvdan.cc/gofumpt@latest
go install github.com/daixiang0/gci@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

# Format code

Simply run:

```bash
make format
```

# Run linters

Several linters are ran using [golangci-lint]https://github.com/golangci/golangci-lint)

For the configuration, see `.golangci.yaml`

```bash
make lint
```