package gen

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/handlename/protoc-gen-oas2connect/internal/oas"
	"github.com/handlename/protoc-gen-oas2connect/internal/proto"
)

func TestBuildTemplateData(t *testing.T) {
	tests := []struct {
		name                 string
		inOasPackageName     string
		inProtoPackagePath   string
		inProtoServiceName   string
		inConnectPackagePath string
		inEndpoints          []Endpoint
		expected             TemplateData
	}{
		{
			name:        "empty",
			inEndpoints: []Endpoint{},
			expected:    TemplateData{},
		},
		{
			name:                 "success",
			inOasPackageName:     "petv1oas",
			inProtoPackagePath:   "example/gen/pet/v1",
			inProtoServiceName:   "Pet",
			inConnectPackagePath: "example/gen/pet/v1/petv1connect",
			inEndpoints: []Endpoint{
				{
					Proto: EndpointProto{
						Service: "Pet",
						Method:  "FindByID",
					},
					Oas: EndpointOas{
						Method: http.MethodGet,
						Path:   "/pet/{petId}",
					},
					Fields: []EndpointField{
						{
							Name: "petId",
							Proto: EndpointProtoField{
								FieldType: proto.FieldTypeInt64,
							},
							Oas: EndpointOasField{
								DataType:   oas.DataTypeNumber,
								DataFormat: oas.DataFormatInt64,
								ParamType:  oas.ParamTypePath,
							},
						},
					},
				},
			},
			expected: TemplateData{
				OasPackageName:     "petv1oas",
				ProtoPackagePath:   "example/gen/pet/v1",
				ConnectPackagePath: "example/gen/pet/v1/petv1connect",
				ProtoServiceName:   "Pet",
				Endpoints: []TemplateEndpointData{
					{
						Path: "/pet/{petId}",
						ProtoMethods: []TemplateProtoMethodData{
							{
								Name:       "FindByID",
								HTTPMethod: http.MethodGet,
								ProtoRequest: TemplateProtoRequestData{
									Name: "FindByID",
									Fields: []TemplateProtoFieldData{
										{
											Name:      "petId",
											GoType:    "int64",
											ParamType: "path",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	g, err := NewGenerator(nil)
	if err != nil {
		t.Fatal(err)
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			res := g.buildTemplateData(
				test.inOasPackageName,
				test.inProtoPackagePath,
				test.inProtoServiceName,
				test.inConnectPackagePath,
				test.inEndpoints,
			)

			if d := cmp.Diff(res, test.expected); d != "" {
				t.Errorf("unexpected result: %v", d)
			}
		})
	}
}
