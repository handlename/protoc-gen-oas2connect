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
	Method      string
	Path        string
	ProtoMethod TemplateProtoMethodData
}

type TemplateProtoMethodData struct {
	Name    string
	Request TemplateProtoRequestData
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
		if d.Endpoints[i].Path == d.Endpoints[j].Path {
			return d.Endpoints[i].Method < d.Endpoints[j].Method
		}

		return d.Endpoints[i].Path < d.Endpoints[j].Path
	})
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
