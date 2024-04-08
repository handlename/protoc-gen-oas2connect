package gen

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTemplateDataFixOrders(t *testing.T) {
	original := &TemplateData{
		Endpoints: []TemplateEndpointData{
			{
				Method: "POST",
				Path:   "/user",
			},
			{
				Method: "GET",
				Path:   "/store",
			},
			{
				Method: "GET",
				Path:   "/pet/{petId}",
			},
			{
				Method: "DELETE",
				Path:   "/pet/{petId}",
			},
			{
				Method: "PATCH",
				Path:   "/pet/{petId}",
			},
		},
	}

	expected := &TemplateData{
		Endpoints: []TemplateEndpointData{
			// Sort by Path, Method
			{
				Method: "DELETE",
				Path:   "/pet/{petId}",
			},
			{
				Method: "GET",
				Path:   "/pet/{petId}",
			},
			{
				Method: "PATCH",
				Path:   "/pet/{petId}",
			},
			{
				Method: "GET",
				Path:   "/store",
			},
			{
				Method: "POST",
				Path:   "/user",
			},
		},
	}

	original.FixOrders()

	if diff := cmp.Diff(original, expected); diff != "" {
		t.Errorf("unexpected result: %v", diff)
	}
}

func TestTemplateExecute(t *testing.T) {
	data := TemplateData{
		OasPackageName:     "petv1oas",
		ProtoServiceName:   "Pet",
		ProtoPackagePath:   "example/gen/pet/v1",
		ConnectPackagePath: "example/gen/pet/v1/petv1connect",
		Endpoints: []TemplateEndpointData{
			{
				Method: "POST",
				Path:   "/pet/{petId}",
				ProtoMethod: TemplateProtoMethodData{
					Name: "UpdatePet",
					Request: TemplateProtoRequestData{
						Name: "",
						Fields: []TemplateProtoFieldData{
							{
								Name:      "name",
								GoType:    "string",
								ParamType: "query",
							},
						},
					},
				},
			},
		},
	}

	var buf bytes.Buffer
	if err := executeTemplate(data, &buf); err != nil {
		t.Fatal(err)
	}
}
