package fiber

import (
	"fmt"
	"testing"
)

func TestgetPathOriginal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		pathOriginal: "test_path",
	}

	result := ctx.getPathOriginal()

	if result != "test_path" {
		t.Errorf("Expected 'test_path', but got '%s'", result)
	}
}
