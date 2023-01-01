format:
	go clean
	go mod tidy
	rm -rf vendor
	go fix
	go mod vendor
	go fmt
	go vet

