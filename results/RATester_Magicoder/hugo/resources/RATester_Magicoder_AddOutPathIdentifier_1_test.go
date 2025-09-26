package resources

import (
	"fmt"
	"testing"
)

func TestAddOutPathIdentifier_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &ResourceTransformationCtx{
		InPath: "/path/to/resource",
	}

	ctx.AddOutPathIdentifier("identifier")

	if ctx.OutPath != "/path/to/resource/identifier" {
		t.Errorf("Expected OutPath to be '/path/to/resource/identifier', but got '%s'", ctx.OutPath)
	}
}
