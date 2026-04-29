.PHONY build test vet :
build_linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/linux/tkl cmd/tkl/main.go

test:
	go test -v -race ./...

vet:
	go vet ./...
