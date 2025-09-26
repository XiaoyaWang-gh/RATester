package fiber

import (
	"fmt"
	"testing"
)

func TestViewBind_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	ctx := &DefaultCtx{}

	vars := Map{
		"key1": "value1",
		"key2": "value2",
	}

	err := ctx.ViewBind(vars)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	value, ok := ctx.viewBindMap.Load("key1")
	if !ok || value != "value1" {
		t.Errorf("Expected 'value1', got %v", value)
	}

	value, ok = ctx.viewBindMap.Load("key2")
	if !ok || value != "value2" {
		t.Errorf("Expected 'value2', got %v", value)
	}
}
