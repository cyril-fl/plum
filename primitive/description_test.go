package primitive

import (
	"reflect"
	"testing"
)

func TestDescribeAndMutators(t *testing.T) {
	descriptible := Describe(TemplateDescriptible{
		Name:        "name",
		Description: "description",
	})

	if descriptible.GetName() != "name" || descriptible.GetDescription() != "description" {
		t.Fatal("Describe did not initialize the name and description")
	}

	descriptible.SetName("updated name")
	descriptible.SetDescription("updated description")

	if descriptible.GetName() != "updated name" || descriptible.GetDescription() != "updated description" {
		t.Fatal("mutators did not update the name and description")
	}
}

func TestDescriptibleHasSerializationTags(t *testing.T) {
	for tagName, want := range map[string]map[string]string{
		"bun": {
			"Name":        "name",
			"Description": "description",
		},
		"bson": {
			"Name":        "name",
			"Description": "description",
		},
		"yaml": {
			"Name":        "name",
			"Description": "description",
		},
		"json": {
			"Name":        "name",
			"Description": "description",
		},
	} {
		t.Run(tagName, func(t *testing.T) {
			assertDescriptibleTags(t, tagName, want)
		})
	}
}

func assertDescriptibleTags(t *testing.T, tagName string, want map[string]string) {
	t.Helper()

	typeOfDescriptible := reflect.TypeOf(Descriptible{})
	for fieldName, expectedTag := range want {
		field, ok := typeOfDescriptible.FieldByName(fieldName)
		if !ok {
			t.Fatalf("Descriptible is missing field %q", fieldName)
		}
		if got := field.Tag.Get(tagName); got != expectedTag {
			t.Errorf("Descriptible.%s %s tag = %q, want %q", fieldName, tagName, got, expectedTag)
		}
	}
}
