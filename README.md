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
# buf.yaml
version: v2
deps:
  - buf.build/googleapis/googleapis
lint:
  use:
    - DEFAULT
  except:
    - FIELD_NOT_REQUIRED
    - PACKAGE_NO_IMPORT_CYCLE
  disallow_comment_ignores: true
breaking:
  use:
    - FILE
  except:
    - EXTENSION_NO_DELETE
    - FIELD_SAME_DEFAULT

```

```yaml
# buf.gen.yaml
version: v2
plugins:
  - remote: buf.build/protocolbuffers/go
    out: gen
    opt:
      - paths=source_relative
  - remote: buf.build/connectrpc/go
    out: gen
    opt:
      - paths=source_relative
  - local: ./protoc-gen-oas2connect
    out: gen
    opt:
      - paths=source_relative
      - connect_package_path=petstore/gen/petstore/v3/petstorev3connect
  - remote: buf.build/grpc-ecosystem/openapiv2
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
