package gen

import (
	"os"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/types/pluginpb"
)

func TestBuildTemplateData(t *testing.T) {
	tests := []struct {
		name                 string
		inTextproto          string
		inProtoPackagePath   string
		inConnectPackagePath string
		expected             TemplateData
	}{
		{
			name:                 "empty",
			inTextproto:          "empty.textproto",
			inProtoPackagePath:   "example/gen/empty/v1",
			inConnectPackagePath: "example/gen/empty/v1/emptyv1connect",
			expected: TemplateData{
				PackageName:        "petstorev3oas",
				ProtoPackagePath:   "example/gen/empty/v1",
				ConnectPackagePath: "example/gen/empty/v1/emptyv1connect",
				Services:           []TemplateServiceData{},
			},
		},
		{
			name:                 "success",
			inTextproto:          "petstore.textproto",
			inProtoPackagePath:   "example/gen/pet/v1",
			inConnectPackagePath: "example/gen/pet/v1/petv1connect",
			expected: TemplateData{
				PackageName:        "petstorev3oas",
				ProtoPackagePath:   "example/gen/pet/v1",
				ConnectPackagePath: "example/gen/pet/v1/petv1connect",
				Services: []TemplateServiceData{
					{
						Name: "Pet",
						Methods: []TemplateMethodData{
							{
								Name:       "FindPet",
								HTTPMethod: "GET",
								HTTPPath:   "/pet/{petId}",
								Request: TemplateRequestData{
									Name: "FindPetRequest",
									Fields: []TemplateFieldData{
										{
											Name:      "PetId",
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

	for _, test := range tests {
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

			res, err := buildTemplateData(p.Files[0], test.inProtoPackagePath, test.inConnectPackagePath)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if d := cmp.Diff(res, &test.expected); d != "" {
				t.Errorf("unexpected result: %v", d)
			}
		})
	}
}
