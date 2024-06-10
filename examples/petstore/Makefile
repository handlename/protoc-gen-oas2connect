PLUGIN := protoc-gen-oas2connect

generate: $(PLUGIN)
	buf generate

$(PLUGIN):
	cd ../.. && go build -o examples/petstore/$@ cmd/protoc-gen-oas2connect/main.go

server:
	go run cmd/server/main.go

clean:
	rm -rf gen
	rm -rf $(PLUGIN)
