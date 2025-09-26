package internal

import (
	"fmt"
	"testing"
)

func TestMarshalJSON_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	namespaces := TemplateFuncsNamespaces{
		{
			Name: "strings",
		},
		{
			Name: "lang",
		},
	}

	expected := `{"strings","lang"}`

	result, err := namespaces.MarshalJSON()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if string(result) != expected {
		t.Fatalf("Expected: %s, got: %s", expected, string(result))
	}
}
