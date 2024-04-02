package gen

import (
	"log/slog"

	"google.golang.org/protobuf/compiler/protogen"
)

type Generator struct {
	files []*protogen.File
}

func NewGenerator(files []*protogen.File) (*Generator, error) {
	return &Generator{files: files}, nil
}

func (g *Generator) Generate() error {
	slog.Error("not implemented yet")
	return nil
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
		protoMethods := []TemplateProtoMethodData{}
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

			protoMethods = append(protoMethods, TemplateProtoMethodData{
				Name:         ep.Proto.Method,
				HTTPMethod:   ep.Oas.Method,
				ProtoRequest: protoRequest,
			})
		}

		data.Endpoints = append(data.Endpoints, TemplateEndpointData{
			Path:         path,
			ProtoMethods: protoMethods,
		})
	}

	data.FixOrders()

	return data
}
