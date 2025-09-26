package codegen

import (
	"fmt"
	"testing"
)

func TestparseSource_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	inspector := &Inspector{
		ProjectRootDir: "testdata",
	}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code panicked with error: %s", r)
		}
	}()

	inspector.parseSource()

	if len(inspector.methodWeight) == 0 {
		t.Error("methodWeight is empty")
	}
}
