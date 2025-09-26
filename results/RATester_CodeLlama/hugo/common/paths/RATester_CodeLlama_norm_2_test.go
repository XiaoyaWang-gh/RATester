package paths

import (
	"fmt"
	"testing"
)

func TestNorm_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		s: "/foo/bar",
	}
	if p.norm("/foo/bar") != "/foo/bar" {
		t.Errorf("norm(%q) = %q, want %q", "/foo/bar", p.norm("/foo/bar"), "/foo/bar")
	}
}
