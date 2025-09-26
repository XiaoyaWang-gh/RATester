package filesystems

import (
	"fmt"
	"testing"
)

func TestIsStatic_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var s SourceFilesystems
	var filename string

	// CONTEXT_0
	filename = "static/js/main.js"
	if s.IsStatic(filename) {
		t.Errorf("IsStatic(%q) = true, want false", filename)
	}

	// CONTEXT_1
	filename = "static/js/main.js"
	if !s.IsStatic(filename) {
		t.Errorf("IsStatic(%q) = false, want true", filename)
	}

	// CONTEXT_2
	filename = "static/js/main.js"
	if s.IsStatic(filename) {
		t.Errorf("IsStatic(%q) = true, want false", filename)
	}
}
