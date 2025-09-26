package internal

import (
	"context"
	"fmt"
	"testing"
)

func TestToJSON_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	namespace := &TemplateFuncsNamespace{
		Name: "test",
		Context: func(ctx context.Context, v ...any) (any, error) {
			return "test context", nil
		},
		MethodMappings: map[string]TemplateFuncMethodMapping{
			"Method1": {
				Method:   "test method 1",
				Aliases:  []string{"alias1", "alias2"},
				Examples: [][2]string{{"input1", "output1"}, {"input2", "output2"}},
			},
			"Method2": {
				Method:   "test method 2",
				Aliases:  []string{"alias3", "alias4"},
				Examples: [][2]string{{"input3", "output3"}, {"input4", "output4"}},
			},
		},
	}

	expectedJSON := `"test": {"Method1":{"Aliases":["alias1","alias2"],"Examples":[["input1","output1"],["input2","output2"]],"Description":"","Args":null},"Method2":{"Aliases":["alias3","alias4"],"Examples":[["input3","output3"],["input4","output4"]],"Description":"","Args":null}}`

	jsonBytes, err := namespace.toJSON(ctx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if string(jsonBytes) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonBytes))
	}
}
