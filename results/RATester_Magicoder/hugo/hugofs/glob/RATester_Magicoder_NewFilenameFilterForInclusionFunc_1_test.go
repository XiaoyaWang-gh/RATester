package glob

import (
	"fmt"
	"testing"
)

func TestNewFilenameFilterForInclusionFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	shouldInclude := func(filename string) bool {
		return true
	}
	filter := NewFilenameFilterForInclusionFunc(shouldInclude)

	if filter == nil {
		t.Error("Expected a non-nil filter, but got nil")
	}

	if filter.shouldInclude == nil {
		t.Error("Expected a non-nil shouldInclude function, but got nil")
	}

	if filter.isWindows {
		t.Error("Expected isWindows to be false, but got true")
	}
}
