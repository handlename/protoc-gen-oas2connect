export LOG_LEVEL := debug

build:
	go build -o protoc-gen-oas2connect cmd/protoc-gen-oas2connect/main.go

test:
	go test -v ./...

clean:
	rm -rf protoc-gen-oas2connect
