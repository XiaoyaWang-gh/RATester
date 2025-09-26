package internal

import (
	"context"
	"fmt"
	"testing"
)

func TesttoJSON_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	namespace := &TemplateFuncsNamespace{
		Name: "test",
		Context: func(ctx context.Context, v ...any) (any, error) {
			return nil, nil
		},
		OnCreated: func(namespaces map[string]any) {},
		MethodMappings: map[string]TemplateFuncMethodMapping{
			"testMethod": {
				Method: func(ctx context.Context, v ...any) (any, error) {
					return nil, nil
				},
				Aliases: []string{"alias1", "alias2"},
				Examples: [][2]string{
					{"input1", "output1"},
					{"input2", "output2"},
				},
			},
		},
	}

	_, err := namespace.toJSON(ctx)
	if err != nil {
		t.Errorf("toJSON() error = %v", err)
		return
	}
}
