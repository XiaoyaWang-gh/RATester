package resources

import (
	"fmt"
	"testing"
)

func TestReplaceOutPathExtension_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &ResourceTransformationCtx{
		InPath: "/path/to/file.txt",
	}

	ctx.ReplaceOutPathExtension(".new")

	if ctx.OutPath != "/path/to/file.new" {
		t.Errorf("Expected OutPath to be '/path/to/file.new', but got '%s'", ctx.OutPath)
	}
}
