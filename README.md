# protoc-gen-oas2connect

> [!CAUTION]
> This product is under development and is not recommended for production use.
> Its interface and behavior may change significantly in future versions.

`protoc-gen-oas2connect` is a code generator plugin for the Protocol Buffers compiler (`protoc`).
It generates Go code that allows you to connect your Application codes using connect-go with REST interfaces defined in OpenAPI Specification (OAS) files.

## Usage

To use `protoc-gen-oas2connect`, you first need to have your Protocol Buffers definitions in `.proto` files. Once you have these, you can run the following command:

```bash
protoc \
    --oas2connect_out=. \
    --oas2connect_opt=connect_package_path=path/to/generated/connect/code \
    --proto_path=./proto \
    yourfile.proto
```

or use [Buf](https://buf.build/) (Recommended)

```yaml
version: v1
breaking:
  use:
    - FILE
lint:
  use:
    - DEFAULT
deps:
  - buf.build/googleapis/googleapis
```

```yaml
# buf.gen.yaml
version: v1
plugins:
  - plugin: go
    out: gen
    opt: paths=source_relative
  - plugin: connect-go
    out: gen
    opt: paths=source_relative
  - plugin: oas2connect
    out: gen
    opt:
      - paths=source_relative
      - connect_package_path=path/to/generated/connect/code
  - plugin: oas
    out: oas
```

This will generate Go code in the current directory, based on the definitions in yourfile.proto.

For more concrete example, please check [examples](./examples) directory.

## Installation

Please download from [Releases](https://github.com/handlename/protoc-gen-oas2connect/releases) and place `protoc-gen-oas2connect` binary in your PATH.

## License

[MIT](./LICENSE)

## Author

@handlename NAGATA Hiroaki https://github.com/handlename
