package hstrings

import (
	"fmt"
	"testing"
)

func TestGetOrCompileRegexp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// CONTEXT_0
	pattern := "pattern"
	// CONTEXT_1
	re, err := GetOrCompileRegexp(pattern)
	if err != nil {
		t.Fatalf("GetOrCompileRegexp(%q) failed: %v", pattern, err)
	}
	// CONTEXT_2
	if re == nil {
		t.Fatalf("GetOrCompileRegexp(%q) returned nil", pattern)
	}
	// CONTEXT_3
	if re.NumSubexp() != 0 {
		t.Fatalf("GetOrCompileRegexp(%q).NumSubexp() = %d, want 0", pattern, re.NumSubexp())
	}
	// CONTEXT_4
	if err != nil {
		t.Fatalf("GetOrCompileRegexp(%q) returned error: %v", pattern, err)
	}
}
