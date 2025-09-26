package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	namespaces := TemplateFuncsNamespaces{
		{
			Name: "strings",
			MethodMappings: map[string]TemplateFuncMethodMapping{
				"title": {
					Examples: [][2]string{{"hello", "Hello"}},
					Aliases:  []string{"t"},
				},
			},
		},
		{
			Name: "lang",
			MethodMappings: map[string]TemplateFuncMethodMapping{
				"upper": {
					Examples: [][2]string{{"hello", "HELLO"}},
					Aliases:  []string{"u"},
				},
			},
		},
	}

	expected := map[string]any{
		"strings": map[string]any{
			"title": map[string]any{
				"Examples": [][2]string{{"hello", "Hello"}},
				"Aliases":  []string{"t"},
			},
		},
		"lang": map[string]any{
			"upper": map[string]any{
				"Examples": [][2]string{{"hello", "HELLO"}},
				"Aliases":  []string{"u"},
			},
		},
	}

	result := namespaces.ToMap()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
