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

	newExt := ".json"
	ctx.ReplaceOutPathExtension(newExt)

	expectedOutPath := "/path/to/file.json"
	if ctx.OutPath != expectedOutPath {
		t.Errorf("Expected OutPath to be %s, but got %s", expectedOutPath, ctx.OutPath)
	}
}
