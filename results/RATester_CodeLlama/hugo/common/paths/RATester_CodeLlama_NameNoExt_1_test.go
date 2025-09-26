package paths

import (
	"fmt"
	"testing"
)

func TestNameNoExt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		s: "a/b/c.d",
	}
	if p.NameNoExt() != "c" {
		t.Errorf("Expected %q, got %q", "c", p.NameNoExt())
	}
}
