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

	ctx := &DefaultCtx{}
	vars := Map{"key": "value"}

	err := ctx.ViewBind(vars)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	value, ok := ctx.viewBindMap.Load("key")
	if !ok {
		t.Errorf("Expected key to be in viewBindMap, but it was not found")
	}

	if value != "value" {
		t.Errorf("Expected value to be 'value', but got: %v", value)
	}
}
