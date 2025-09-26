package fiber

import (
	"fmt"
	"testing"
)

func TestgetTreePath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		treePath: "test_path",
	}

	result := ctx.getTreePath()

	if result != "test_path" {
		t.Errorf("Expected 'test_path', but got '%s'", result)
	}
}
