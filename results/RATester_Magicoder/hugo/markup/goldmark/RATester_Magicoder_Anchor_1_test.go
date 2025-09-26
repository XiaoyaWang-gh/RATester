package goldmark

import (
	"fmt"
	"testing"
)

func TestAnchor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := headingContext{
		anchor: "test-anchor",
	}

	result := ctx.Anchor()

	if result != "test-anchor" {
		t.Errorf("Expected 'test-anchor', but got '%s'", result)
	}
}
