package passthrough

import (
	"fmt"
	"testing"
)

func TestInner_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &passthroughContext{
		typ:   "inner",
		inner: "test",
	}

	result := ctx.Inner()

	if result != "test" {
		t.Errorf("Expected 'test', but got '%s'", result)
	}
}
