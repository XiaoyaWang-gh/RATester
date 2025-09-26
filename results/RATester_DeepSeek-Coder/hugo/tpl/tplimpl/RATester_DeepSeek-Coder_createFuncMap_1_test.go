package tplimpl

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestCreateFuncMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &deps.Deps{
		OverloadedTemplateFuncs: map[string]any{
			"testFunc": func(args ...any) (any, error) {
				return "testFunc", nil
			},
		},
	}

	funcMap := createFuncMap(d)

	if _, exists := funcMap["testFunc"]; !exists {
		t.Errorf("Expected 'testFunc' to be in the funcMap")
	}

	testFunc, err := funcMap["testFunc"].(func(args ...any) (any, error))("testArg")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if testFunc != "testFunc" {
		t.Errorf("Expected 'testFunc', got %v", testFunc)
	}
}
