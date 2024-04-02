package gen

import (
	_ "embed"
)

//go:embed template/service.go.tmpl
var ServiceTmpl []byte

type TemplateData struct {
	OasPackageName     string
	ProtoPackagePath   string
	ProtoServiceName   string
	ConnectPackagePath string
	Endpoints          []TemplateEndpointData
}

type TemplateEndpointData struct {
	Path         string
	ProtoMethods []TemplateProtoMethodData
}

type TemplateProtoMethodData struct {
	Name         string
	HTTPMethod   string
	ProtoRequest TemplateProtoRequestData
}

type TemplateProtoRequestData struct {
	Name   string
	Fields []TemplateProtoFieldData
}

type TemplateProtoFieldData struct {
	Name      string
	GoType    string
	ParamType string // "query" or "path"
}
