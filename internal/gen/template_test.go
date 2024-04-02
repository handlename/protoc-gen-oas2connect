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
				Path:         "/user",
				ProtoMethods: []TemplateProtoMethodData{},
			},
			{
				Path:         "/store",
				ProtoMethods: []TemplateProtoMethodData{},
			},
			{
				Path: "/pet/{petId}",
				ProtoMethods: []TemplateProtoMethodData{
					{
						Name:       "FindPetByID",
						HTTPMethod: "GET",
						ProtoRequest: TemplateProtoRequestData{
							Name:   "",
							Fields: []TemplateProtoFieldData{},
						},
					},
					{
						Name:       "DeletePet",
						HTTPMethod: "DELETE",
						ProtoRequest: TemplateProtoRequestData{
							Name:   "",
							Fields: []TemplateProtoFieldData{},
						},
					},
					{
						Name:       "UpdatePet",
						HTTPMethod: "POST",
						ProtoRequest: TemplateProtoRequestData{
							Name: "",
							Fields: []TemplateProtoFieldData{
								{
									Name:      "petId",
									GoType:    "int64",
									ParamType: "path",
								},
								{
									Name:      "name",
									GoType:    "string",
									ParamType: "query",
								},
								{
									Name:      "status",
									GoType:    "string",
									ParamType: "query",
								},
							},
						},
					},
				},
			},
		},
	}

	expected := &TemplateData{
		Endpoints: []TemplateEndpointData{
			// Sort by Path
			{
				Path: "/pet/{petId}",
				ProtoMethods: []TemplateProtoMethodData{
					// Sort by HTTPMethod
					{
						Name:       "DeletePet",
						HTTPMethod: "DELETE",
						ProtoRequest: TemplateProtoRequestData{
							Name:   "",
							Fields: []TemplateProtoFieldData{},
						},
					},
					{
						Name:       "FindPetByID",
						HTTPMethod: "GET",
						ProtoRequest: TemplateProtoRequestData{
							Name:   "",
							Fields: []TemplateProtoFieldData{},
						},
					},
					{
						Name:       "UpdatePet",
						HTTPMethod: "POST",
						ProtoRequest: TemplateProtoRequestData{
							Name: "",
							Fields: []TemplateProtoFieldData{
								// Sort by Name
								{
									Name:      "name",
									GoType:    "string",
									ParamType: "query",
								},
								{
									Name:      "petId",
									GoType:    "int64",
									ParamType: "path",
								},
								{
									Name:      "status",
									GoType:    "string",
									ParamType: "query",
								},
							},
						},
					},
				},
			},
			{
				Path:         "/store",
				ProtoMethods: []TemplateProtoMethodData{},
			},
			{
				Path:         "/user",
				ProtoMethods: []TemplateProtoMethodData{},
			},
		},
	}

	original.FixOrders()

	if diff := cmp.Diff(original, expected); diff != "" {
		t.Errorf("unexpected result: %v", diff)
	}
}

func TestExecuteTemplate(t *testing.T) {
	data := TemplateData{
		OasPackageName:     "petv1oas",
		ProtoServiceName:   "Pet",
		ProtoPackagePath:   "example/gen/pet/v1",
		ConnectPackagePath: "example/gen/pet/v1/petv1connect",
		Endpoints: []TemplateEndpointData{
			{
				Path: "/pet/{petId}",
				ProtoMethods: []TemplateProtoMethodData{
					{
						Name:       "UpdatePet",
						HTTPMethod: "POST",
						ProtoRequest: TemplateProtoRequestData{
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
		},
	}

	var buf bytes.Buffer
	if err := executeTemplate(data, &buf); err != nil {
		t.Fatal(err)
	}
}
