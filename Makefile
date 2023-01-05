PACKAGES = cmd core processors utils worker

all:
	@mkdir -p bin
	GOOS=windows GOARCH=amd64 go build -o bin/forensibus_x64.exe main.go
	CGO_ENABLED=0 go build -ldflags="-extldflags=-static" -o bin/forensibus main.go
	sudo chroot . ./bin/forensibus

format:
	go mod tidy
	rm -rf vendor
	go fix
	go mod vendor
	gofmt -w -s  $(PACKAGES)
	# go vet $(PACKAGES)
	go clean

test:
	@echo "****** Running tests ******"
	@go test -cover  ./...  -coverprofile=coverage.out
	@echo "\n****** Generating coverage.out ******"
	@go tool cover -html=coverage.out
	@echo "\n****** Tests coverage ******"
	@go tool cover -func=coverage.out

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/worker.proto

.PHONY: proto