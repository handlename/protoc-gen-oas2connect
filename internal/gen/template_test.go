package gen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTemplateDataFixOrders(t *testing.T) {
	original := &TemplateData{
		Services: []TemplateServiceData{
			{
				Name: "Store",
				Methods: []TemplateMethodData{
					{Name: "FindStore"},
					{Name: "UpdateStore"},
				},
			},
			{
				Name: "Pet",
				Methods: []TemplateMethodData{
					{Name: "FindPet"},
					{Name: "AddPet"},
					{Name: "DeletePet"},
				},
			},
		},
	}

	expected := &TemplateData{
		Services: []TemplateServiceData{
			{
				Name: "Pet",
				Methods: []TemplateMethodData{
					{Name: "AddPet"},
					{Name: "DeletePet"},
					{Name: "FindPet"},
				},
			},
			{
				Name: "Store",
				Methods: []TemplateMethodData{
					{Name: "FindStore"},
					{Name: "UpdateStore"},
				},
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
		PackageName:        "",
		ProtoPackagePath:   "",
		ConnectPackagePath: "",
		Services: []TemplateServiceData{
			{
				Name: "",
				Methods: []TemplateMethodData{
					{
						Name:       "",
						HTTPMethod: "",
						HTTPPath:   "",
						Request: TemplateRequestData{
							Name:       "",
							ExpectBody: false,
							Fields: []TemplateFieldData{
								{
									Name:      "",
									GoType:    "",
									ParamType: "",
								},
							},
						},
					},
				},
			},
		},
	}

	if _, err := executeTemplate("Service", data); err != nil {
		t.Fatal(err)
	}
}
