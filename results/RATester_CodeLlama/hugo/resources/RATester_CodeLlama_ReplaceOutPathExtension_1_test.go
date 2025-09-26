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
		InPath: "foo/bar.txt",
	}
	ctx.ReplaceOutPathExtension("newExt")
	if ctx.OutPath != "foo/bar.newExt" {
		t.Errorf("Expected %q, got %q", "foo/bar.newExt", ctx.OutPath)
	}
}
