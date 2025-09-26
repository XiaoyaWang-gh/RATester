package template

import (
	"fmt"
	"testing"
)

func TestcreateValueFuncs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	funcMap := FuncMap{
		"testFunc": func() {},
	}

	result := createValueFuncs(funcMap)

	if len(result) != 1 {
		t.Errorf("Expected 1 function, got %d", len(result))
	}

	_, ok := result["testFunc"]
	if !ok {
		t.Error("Expected function 'testFunc' not found in result")
	}
}
