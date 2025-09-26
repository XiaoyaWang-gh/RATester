package codegen

import (
	"fmt"
	"testing"
)

func TestParseSource_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	inspector := &Inspector{
		ProjectRootDir: "testdata",
		methodWeight:   make(map[string]map[string]int),
	}

	inspector.parseSource()

	if len(inspector.methodWeight) == 0 {
		t.Errorf("Expected methodWeight to be populated, got %v", inspector.methodWeight)
	}
}
