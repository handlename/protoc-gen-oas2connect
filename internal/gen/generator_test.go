package gen

import (
	"os"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/handlename/protoc-gen-oas2connect/internal/oas"
	"github.com/handlename/protoc-gen-oas2connect/internal/proto"
)

func TestBuildEndpoints(t *testing.T) {
	test := []struct {
		name        string
		inTextproto string
		expected    []Endpoint
	}{
		{
			name:        "empty",
			inTextproto: "empty.textproto",
			expected:    []Endpoint{},
		},
		{
			name:        "success",
			inTextproto: "petstore.textproto",
			expected: []Endpoint{
				// TODO: POST body
				// {
				// 	Proto:  EndpointProto{
				// 		Service: "Pet",
				// 		Method:  "AddNewPet",
				// 	},
				// 	Oas:    EndpointOas{
				// 		Method: "POST",
				// 		Path:   "/pet",
				// 	},
				// },
				{
					Proto: EndpointProto{
						Service: "Pet",
						Method:  "FindPetByID",
					},
					Oas: EndpointOas{
						Method: "GET",
						Path:   "/pet/{petId}",
					},
					Fields: []EndpointField{
						{
							Name: "PetId",
							Proto: EndpointProtoField{
								FieldType: proto.FieldTypeInt64,
							},
							Oas: EndpointOasField{
								DataType:  oas.DataTypeInteger,
								ParamType: oas.ParamTypePath,
							},
						},
					},
				},
			},
		},
	}

	for _, test := range test {
		test := test

		t.Run(test.name, func(t *testing.T) {
			textproto, err := os.ReadFile(path.Join("_testdata", test.inTextproto))
			if err != nil {
				t.Fatalf("failed to read textproto: %v", err)
			}

			req := new(pluginpb.CodeGeneratorRequest)
			err = prototext.Unmarshal(textproto, req)
			if err != nil {
				t.Fatalf("failed to unmarshal textproto: %v", err)
			}

			opts := protogen.Options{}
			p, err := opts.New(req)
			if err != nil {
				t.Fatalf("failed to create plugin: %v", err)
			}

			for _, f := range p.Files {
				f.Generate = true
			}

			g, err := NewGenerator(p.Files)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(g.buildEndpoints(), test.expected); diff != "" {
				t.Errorf("unexpected result: %v", diff)
			}
		})
	}

}

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
						Method: "GET",
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
						Method: "GET",
						Path:   "/pet/{petId}",
						ProtoMethod: TemplateProtoMethodData{
							Name: "FindByID",
							Request: TemplateProtoRequestData{
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
