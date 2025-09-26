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

	identifier := "test-identifier"
	ctx.AddOutPathIdentifier(identifier)

	expectedOutPath := ctx.addPathIdentifier(ctx.InPath, identifier)
	if ctx.OutPath != expectedOutPath {
		t.Errorf("Expected OutPath to be %s, but got %s", expectedOutPath, ctx.OutPath)
	}
}
