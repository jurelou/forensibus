PACKAGES = cmd core processors utils worker


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
	go install github.com/golang/protobuf/protoc-gen-go
	protoc  --proto_path proto --go_out=plugins=grpc:proto/ proto/kv.proto


.PHONY: proto