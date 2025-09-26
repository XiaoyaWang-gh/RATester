package paths

import (
	"fmt"
	"testing"
)

func TestTrimLeadingSlash_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := &Path{s: "/test/path"}
	result := path.TrimLeadingSlash()

	if result.s != "test/path" {
		t.Errorf("Expected '/test/path', got '%s'", result.s)
	}

	if !result.trimLeadingSlash {
		t.Error("Expected trimLeadingSlash to be true")
	}
}
