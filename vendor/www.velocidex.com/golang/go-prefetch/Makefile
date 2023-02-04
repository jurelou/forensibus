all:
	go build -ldflags="-s -w"  -o prefetch cmd/*.go

gen:
	binparsegen conversion.spec.yaml  > prefetch_gen.go

test:
	go test ./...

windows:
	GOOS=windows GOARCH=amd64 go build  -ldflags="-s -w" -o prefetch.exe cmd/*.go
