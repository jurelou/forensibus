PACKAGES = cmd core processors utils worker

LDFLAGS = -ldflags="-s -w"
GO_BUILD_TAGS = -tags yara_static
MODULES = $(shell go list -f '{{.Dir}}' -v ./...)
BIN_FOLDER = ./bin

all: windows linux

windows:
	mkdir -p $(BIN_FOLDER)
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build $(GO_BUILD_TAGS) -o $(BIN_FOLDER)/forensibus_windows_amd64.exe main.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc go build -o $(BIN_FOLDER)/forensibus_windows_x86.exe main.go

# darwin:
# 	mkdir -p $(BIN_FOLDER)
# 	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) $(GO_BUILD_TAGS) -o ${BIN_FOLDER}/forensibus_darwin_amd64 main.go

linux:
	mkdir -p $(BIN_FOLDER)
	CC=/usr/local/musl/bin/musl-gcc GOOS=linux GOARCH=amd64 go build $(GO_BUILD_TAGS) -ldflags='-s -w -linkmode external -extldflags "-static"' -o $(BIN_FOLDER)/forensibus_linux_amd64 main.go

install:
	go install mvdan.cc/gofumpt@latest
	go install github.com/daixiang0/gci@latest
	go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest

docs:
	cd docs; npm run start

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

splunk:
	docker-compose -f docker/docker-compose.yml up -d

resplunk:
	docker-compose -f docker/docker-compose.yml down --volumes --rmi local
	docker-compose -f docker/docker-compose.yml up -d --build --force-recreate
	
test:
	@echo "****** Running tests ******"
	@go test -cover  ./...  -coverprofile=coverage.out
	@echo "\n****** Generating coverage.out ******"
	@go tool cover -html=coverage.out
	@echo "\n****** Tests coverage ******"
	@go tool cover -func=coverage.out

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/worker/worker.proto

changelog:
	git-chglog -o CHANGELOG.md

release: vendor format all

.PHONY: proto vendor splunk resplunk docs