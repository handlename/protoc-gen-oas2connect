package gen

import (
	_ "embed"
	"fmt"
	"io"
	"sort"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed template/convert.go.tmpl
var ConvertTmpl []byte

//go:embed template/service.go.tmpl
var ServiceTmpl []byte

var Templates = map[string][]byte{
	"Convert": ConvertTmpl,
	"Service": ServiceTmpl,
}

type TemplateData struct {
	PackageName        string
	ProtoPackagePath   string
	ConnectPackagePath string
	Services           []TemplateServiceData
}

type TemplateServiceData struct {
	Name string
	// Endpoints []TemplateEndpointData
	Methods []TemplateMethodData
}

type TemplateMethodData struct {
	Name       string
	HTTPMethod string
	HTTPPath   string
	Request    TemplateRequestData
}

// TempalteRequestData is a data for proto request message
type TemplateRequestData struct {
	Name       string
	ExpectBody bool
	Fields     []TemplateFieldData
}

type TemplateFieldData struct {
	// Name is name used in proto file
	Name string

	// GoName is name used in Go code
	GoName string

	// GoType is type used in Go code
	GoType string

	// ParamType is where the field is located
	// "query", "path" or "body"
	ParamType string

	// Repeated is true if the field is allowed multiple values
	Repeated bool

	// Optional is true if the field is optional
	Optional bool
}

// FixOrders fixes the order of Services by its Name
func (d *TemplateData) FixOrders() {
	sort.SliceStable(d.Services, func(i, j int) bool {
		return d.Services[i].Name < d.Services[j].Name
	})

	for _, service := range d.Services {
		service.FixOrders()
	}
}

// FixOrders fixes the order of Methods by its Name
func (d *TemplateServiceData) FixOrders() {
	sort.SliceStable(d.Methods, func(i, j int) bool {
		return d.Methods[i].Name < d.Methods[j].Name
	})

	for _, method := range d.Methods {
		method.Request.FixOrders()
	}
}

// FixOrders fixes the order of Fields by its Name
func (d *TemplateRequestData) FixOrders() {
	sort.SliceStable(d.Fields, func(i, j int) bool {
		return d.Fields[i].Name < d.Fields[j].Name
	})
}

func executeTemplate(name string, data TemplateData, out io.Writer) error {
	tmpl := template.New(name)
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

	tmplSrc, ok := Templates[name]
	if !ok {
		return fmt.Errorf("template %s not found", name)
	}

	tmpl, err := tmpl.Parse(string(tmplSrc))
	if err != nil {
		return err
	}

	if err := tmpl.ExecuteTemplate(out, name, data); err != nil {
		return err
	}

	return nil
}
