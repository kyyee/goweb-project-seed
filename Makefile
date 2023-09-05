.PHONY: build
build:
	echo "building"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go21 build -mod=vendor -ldflags="-s -w" -o ./ ./cmd/server/server.go