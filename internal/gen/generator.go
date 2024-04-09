package gen

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/handlename/protoc-gen-oas2connect/internal/proto"
	"github.com/samber/lo"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	pt "google.golang.org/protobuf/proto"
)

var paramsInPathRgx = regexp.MustCompile(`\{([^}]+)\}`)

// Generate generates a go code to connect OpenAPI interface with connect service methods.
func Generate(file *protogen.File, protoPackagePath, connectPackagePath string, out io.Writer) error {
	data, err := buildTemplateData(file, protoPackagePath, connectPackagePath)
	if err != nil {
		return err
	}

	return GenerateWithData(data, out)
}

func GenerateWithData(data *TemplateData, out io.Writer) error {
	return executeTemplate("Service", *data, out)
}

func buildTemplateData(file *protogen.File, protoPackagePath, connectPackagePath string) (*TemplateData, error) {
	data := TemplateData{
		PackageName:        string(file.GoPackageName),
		ProtoPackagePath:   protoPackagePath,
		ConnectPackagePath: connectPackagePath,
		Services:           []TemplateServiceData{},
	}

	for _, service := range file.Services {
		serviceData := TemplateServiceData{
			Name: service.GoName,
		}

		for _, method := range service.Methods {
			methodData, err := buildTemplateMethodData(method)
			if err != nil {
				return nil, err
			}

			serviceData.Methods = append(serviceData.Methods, *methodData)
		}

		data.Services = append(data.Services, serviceData)
	}

	return &data, nil
}

func buildTemplateMethodData(method *protogen.Method) (*TemplateMethodData, error) {
	data := TemplateMethodData{
		Name: method.GoName,
		Request: TemplateRequestData{
			Name:   method.Input.GoIdent.GoName,
			Fields: []TemplateFieldData{},
		},
	}

	rule, ok := pt.GetExtension(method.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)
	if !ok {
		return nil, fmt.Errorf("no http rule for method %s", method.GoName)
	}

	switch rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		data.HTTPMethod = http.MethodGet
	case *annotations.HttpRule_Post:
		data.HTTPMethod = http.MethodPost
	case *annotations.HttpRule_Put:
		data.HTTPMethod = http.MethodPut
	case *annotations.HttpRule_Delete:
		data.HTTPMethod = http.MethodDelete
	case *annotations.HttpRule_Patch:
		data.HTTPMethod = http.MethodPatch
	}

	switch p := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		data.HTTPPath = p.Get
	case *annotations.HttpRule_Post:
		data.HTTPPath = p.Post
	case *annotations.HttpRule_Put:
		data.HTTPPath = p.Put
	case *annotations.HttpRule_Delete:
		data.HTTPPath = p.Delete
	case *annotations.HttpRule_Patch:
		data.HTTPPath = p.Patch
	}

	if rule.Body == "*" {
		data.Request.ExpectBody = true
	}

	paramsInPath := (func() []string {
		m := paramsInPathRgx.FindAllStringSubmatch(data.HTTPPath, -1)
		return lo.Map(m, func(mm []string, i int) string {
			return mm[1]
		})
	})()

	for _, f := range method.Input.Fields {
		paramType := "query"
		if lo.Contains(paramsInPath, f.Desc.TextName()) {
			paramType = "path"
		} else if data.Request.ExpectBody {
			continue
		}

		data.Request.Fields = append(data.Request.Fields, TemplateFieldData{
			Name:      f.GoName,
			GoType:    proto.NewFieldType(f.Desc.Kind()).ToGoTypeName(),
			ParamType: paramType,
			Repeated:  f.Desc.IsList(),
		})
	}

	return &data, nil
}

func GenerateOther(name, packageName string, out io.Writer) error {
	data := TemplateData{
		PackageName: packageName,
	}

	return executeTemplate(name, data, out)
}
