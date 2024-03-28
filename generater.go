package o2c

import (
	_ "embed"
)

//go:embed template/service.go.tmpl
var ServiceTmpl []byte

type TemplateData struct {
	OasPackage     TemplatePackageData
	ProtoPackage   TemplatePackageData
	ConnectPackage TemplatePackageData
	ProtoService   TemplateProtoServiceData
	Endpoints      []TemplateEndpointData
}

type TemplatePackageData struct {
	Name string
	Path string
}

type TemplateProtoServiceData struct {
	Name string
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
