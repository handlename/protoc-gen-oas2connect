package gen

import (
	_ "embed"
	"io"
	"sort"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

func executeTemplate(data TemplateData, out io.Writer) error {
	tmpl := template.New("service")
	tmpl = tmpl.Funcs(template.FuncMap{
		"PathToFuncName": func(s string) string {
			s = strings.ReplaceAll(s, "/", "_")
			s = strings.ReplaceAll(s, "{", "")
			s = strings.ReplaceAll(s, "}", "")
			return s
		},
		"ToCamel": func(s string) string {
			return cases.Title(language.English).String(s)
		},
	})

	tmpl, err := tmpl.Parse(string(ServiceTmpl))
	if err != nil {
		return err
	}

	if err := tmpl.ExecuteTemplate(out, "Service", data); err != nil {
		return err
	}

	return nil
}
