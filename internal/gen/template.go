package gen

import (
	_ "embed"
	"sort"
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

// FixOrders fixes the order of Endpoints by Path
func (d *TemplateData) FixOrders() {
	sort.SliceStable(d.Endpoints, func(i, j int) bool {
		return d.Endpoints[i].Path < d.Endpoints[j].Path
	})

	for _, ep := range d.Endpoints {
		ep.FixOrders()
	}
}

// FixOrders fixes the order of ProtoMethods by HTTPMethod
func (d *TemplateEndpointData) FixOrders() {
	sort.SliceStable(d.ProtoMethods, func(i, j int) bool {
		return d.ProtoMethods[i].HTTPMethod < d.ProtoMethods[j].HTTPMethod
	})

	for _, m := range d.ProtoMethods {
		m.ProtoRequest.FixOrders()
	}
}

// FixOrders fixes the order of Fields by Name
func (d *TemplateProtoRequestData) FixOrders() {
	sort.SliceStable(d.Fields, func(i, j int) bool {
		return d.Fields[i].Name < d.Fields[j].Name
	})
}
