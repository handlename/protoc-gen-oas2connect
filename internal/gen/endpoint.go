package gen

import (
	"github.com/handlename/protoc-gen-oas2connect/internal/oas"
	"github.com/handlename/protoc-gen-oas2connect/internal/proto"
)

type Endpoint struct {
	Proto  EndpointProto
	Oas    EndpointOas
	Fields []EndpointField
}

type EndpointField struct {
	Name  string
	Proto EndpointProtoField
	Oas   EndpointOasField
}

type EndpointProto struct {
	Service string
	Method  string
}

type EndpointOas struct {
	Method string // http method
	Path   string
}

type EndpointProtoField struct {
	FieldType proto.FieldType
}

type EndpointOasField struct {
	DataType   oas.DataType
	DataFormat oas.DataFormat
	ParamType  oas.ParamType
}
