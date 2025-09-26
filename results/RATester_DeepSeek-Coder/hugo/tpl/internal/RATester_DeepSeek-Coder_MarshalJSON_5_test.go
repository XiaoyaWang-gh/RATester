package internal

import (
	"context"
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
		&TemplateFuncsNamespace{
			Name: "test",
			Context: func(ctx context.Context, v ...any) (any, error) {
				return "test", nil
			},
			OnCreated: func(namespaces map[string]any) {
				// do something
			},
			MethodMappings: map[string]TemplateFuncMethodMapping{
				"test": {
					Method: func(ctx context.Context, v ...any) (any, error) {
						return "test", nil
					},
					Aliases:  []string{"t"},
					Examples: [][2]string{{"input", "expected"}},
				},
			},
		},
	}

	expected := `{"test":{"test":"test","t":["input","expected"]}}`

	result, err := namespaces.MarshalJSON()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if string(result) != expected {
		t.Errorf("Expected %s, got %s", expected, string(result))
	}
}
