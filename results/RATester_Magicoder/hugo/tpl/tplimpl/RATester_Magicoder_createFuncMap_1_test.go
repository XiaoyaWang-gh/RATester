package tplimpl

import (
	"context"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestcreateFuncMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &deps.Deps{
		OverloadedTemplateFuncs: map[string]any{
			"testFunc": func(ctx context.Context) (any, error) {
				return "testValue", nil
			},
		},
	}

	funcMap := createFuncMap(d)

	if _, exists := funcMap["testFunc"]; !exists {
		t.Error("Expected testFunc to be in the funcMap")
	}

	if _, exists := funcMap["testValue"]; exists {
		t.Error("Expected testValue to not be in the funcMap")
	}
}
