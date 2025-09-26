package blockquotes

import (
	"fmt"
	"testing"
)

func TestType_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &blockquoteContext{
		typ: "testType",
	}

	expected := "testType"
	actual := ctx.Type()

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
