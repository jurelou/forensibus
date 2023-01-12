PACKAGES = cmd core processors utils worker

UNIX_LDFLAGS = -ldflags="-s -w -linkmode external -extldflags=-static"

MODULES = $(shell go list -f '{{.Dir}}' -v ./...)

all:
	@mkdir -p bin
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o ./bin/forensibus_windows_amd64.exe main.go
	GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc go build -o ./bin/forensibus_windows_x86.exe main.go

	# GOOS=darwin GOARCH=amd64 go build $(UNIX_LDFLAGS) -o ./bin/forensibus_darwin_amd64 main.go
	# GOOS=darwin GOARCH=arm64 go build -ldflags="-extldflags=-static" -o ./bin/forensibus_darwin_arm64 main.go

	GOOS=linux GOARCH=amd64 go build $(UNIX_LDFLAGS) -o ./bin/forensibus_linux_amd64 main.go 2>&1 | grep -vi warning
	# GOOS=linux GOARCH=386 go build $(UNIX_LDFLAGS)  -o ./bin/forensibus_linux_x86 main.go 2>&1 | grep -vi warning

	#sudo chroot . ./forensibus

install:
	go install mvdan.cc/gofumpt@latest
	go install github.com/daixiang0/gci@latest


lint:
	golangci-lint run

format:
	go fix
	gofumpt -l -w .
	gci write --skip-generated --section "standard,default,prefix(github.com/jurelou/forensibus),blank,dot" $(MODULES)
	go clean

vendor:
	rm -rf vendor
	go mod tidy
	go mod vendor
	go clean

test:
	@echo "****** Running tests ******"
	@go test -cover  ./...  -coverprofile=coverage.out
	@echo "\n****** Generating coverage.out ******"
	@go tool cover -html=coverage.out
	@echo "\n****** Tests coverage ******"
	@go tool cover -func=coverage.out

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/worker/worker.proto

release: vendor format all

.PHONY: proto vendor