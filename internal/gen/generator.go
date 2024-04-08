package gen

import (
	"bytes"
	"log/slog"
	"net/http"
	"regexp"

	"github.com/handlename/protoc-gen-oas2connect/internal/oas"
	"github.com/handlename/protoc-gen-oas2connect/internal/proto"
	"github.com/samber/lo"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	pt "google.golang.org/protobuf/proto"
)

var paramsInPathRgx = regexp.MustCompile(`\{([^}]+)\}`)

type Generator struct {
	files []*protogen.File
}

type GenerateCallback func(filename string, content []byte) error

func NewGenerator(files []*protogen.File) (*Generator, error) {
	return &Generator{files: files}, nil
}

func (g *Generator) Generate(cb GenerateCallback) error {
	// build Endpoints
	slog.Debug("building Endpoints")
	endpoints := g.buildEndpoints()
	slog.Debug("built Endpoints", slog.Int("count", len(endpoints)))

	services := map[string][]Endpoint{}
	for _, ep := range endpoints {
		services[ep.Proto.Service] = append(services[ep.Proto.Service], ep)
	}

	for serviceName, eps := range services {
		// build TemplateData per each proto Service
		slog.Debug("processing to generate code", slog.String("service", serviceName))

		data := g.buildTemplateData(
			"oas",   // TODO
			"proto", // TODO
			serviceName,
			"connect",
			eps,
		)

		// render template
		slog.Debug("rendering template", slog.String("service", serviceName))

		var buf bytes.Buffer
		if err := executeTemplate(data, &buf); err != nil {
			slog.Error("failed to execute template", slog.String("err", err.Error()), slog.String("service", serviceName))
			return err
		}

		filename := serviceName + "oas2connect.go"
		slog.Debug("generating code", slog.String("service", serviceName), slog.String("filename", filename))
		if err := cb(filename, buf.Bytes()); err != nil {
			slog.Error("failed to call callback", slog.String("err", err.Error()), slog.String("service", serviceName))
			return err
		}

		slog.Debug("generated code", slog.String("service", serviceName))
	}

	return nil
}

func (g *Generator) buildEndpoints() []Endpoint {
	eps := []Endpoint{}

	for _, f := range g.files {
		for _, s := range f.Services {
			for _, m := range s.Methods {
				slog.Debug(
					"processing to build Endpoint",
					slog.String("service", s.GoName),
					slog.String("method", m.GoName),
				)

				ep := g.buildEndpoint(s, m)
				if ep != nil {
					eps = append(eps, *ep)
				}
			}
		}
	}

	return eps
}

func (g *Generator) buildEndpoint(service *protogen.Service, method *protogen.Method) *Endpoint {
	ep := &Endpoint{}

	// build EndpointProto

	ep.Proto = EndpointProto{
		Service: service.GoName,
		Method:  method.GoName,
	}

	// build EndpointOas

	rule, ok := pt.GetExtension(method.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)
	if !ok {
		slog.Debug(
			"no http rule",
			slog.String("service", service.GoName),
			slog.String("method", method.GoName),
		)
		return nil
	}

	ep.Oas = EndpointOas{}

	switch rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		ep.Oas.Method = http.MethodGet
	case *annotations.HttpRule_Post:
		ep.Oas.Method = http.MethodPost
	case *annotations.HttpRule_Put:
		ep.Oas.Method = http.MethodPut
	case *annotations.HttpRule_Delete:
		ep.Oas.Method = http.MethodDelete
	case *annotations.HttpRule_Patch:
		ep.Oas.Method = http.MethodPatch
	}

	switch p := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		ep.Oas.Path = p.Get
	case *annotations.HttpRule_Post:
		ep.Oas.Path = p.Post
	case *annotations.HttpRule_Put:
		ep.Oas.Path = p.Put
	case *annotations.HttpRule_Delete:
		ep.Oas.Path = p.Delete
	case *annotations.HttpRule_Patch:
		ep.Oas.Path = p.Patch
	}

	// build EndpointFields

	ep.Fields = []EndpointField{}

	paramsInPath := (func() []string {
		m := paramsInPathRgx.FindAllStringSubmatch(ep.Oas.Path, -1)
		return lo.Map(m, func(mm []string, i int) string {
			return mm[1]
		})
	})()

	for _, f := range method.Input.Fields {
		paramType := "query"
		if lo.Contains(paramsInPath, f.Desc.TextName()) {
			paramType = "path"
		}

		ep.Fields = append(ep.Fields, EndpointField{
			Name: f.GoName,
			Proto: EndpointProtoField{
				FieldType: proto.NewFieldType(f.Desc.Kind()),
			},
			Oas: EndpointOasField{
				DataType:   toOasDataType(proto.NewFieldType(f.Desc.Kind())),
				DataFormat: "", // not supported yet
				ParamType:  oas.ParamType(paramType),
			},
		})
	}

	return ep
}

func (g *Generator) buildTemplateData(oasPackageName, protoPackagePath, protoServiceName, connectPackagePath string, endpoints []Endpoint) TemplateData {
	data := TemplateData{
		OasPackageName:     oasPackageName,
		ProtoPackagePath:   protoPackagePath,
		ConnectPackagePath: connectPackagePath,
		ProtoServiceName:   protoServiceName,
	}

	byPath := map[string][]Endpoint{}

	for _, ep := range endpoints {
		byPath[ep.Oas.Path] = append(byPath[ep.Oas.Path], ep)
	}

	for path, eps := range byPath {
		for _, ep := range eps {
			fields := []TemplateProtoFieldData{}
			for _, f := range ep.Fields {
				fields = append(fields, TemplateProtoFieldData{
					Name:      f.Name,
					GoType:    f.Proto.FieldType.ToGoTypeName(),
					ParamType: string(f.Oas.ParamType),
				})
			}

			protoRequest := TemplateProtoRequestData{
				Name:   ep.Proto.Method,
				Fields: fields,
			}

			protoMethod := TemplateProtoMethodData{
				Name:    ep.Proto.Method,
				Request: protoRequest,
			}

			data.Endpoints = append(data.Endpoints, TemplateEndpointData{
				Method:      ep.Oas.Method,
				Path:        path,
				ProtoMethod: protoMethod,
			})
		}
	}

	data.FixOrders()

	return data
}
