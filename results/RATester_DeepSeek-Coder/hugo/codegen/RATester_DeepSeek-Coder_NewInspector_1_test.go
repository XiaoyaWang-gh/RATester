package codegen

import (
	"fmt"
	"testing"
)

func TestNewInspector_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	root := "/path/to/root"
	inspector := NewInspector(root)

	if inspector.ProjectRootDir != root {
		t.Errorf("Expected ProjectRootDir to be %s, got %s", root, inspector.ProjectRootDir)
	}
}
