package page

import (
	"fmt"
	"testing"
)

func TestPathDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pagePathBuilder{}
	p.els = []string{"a", "b", "c"}
	p.prefixPath = "prefix"
	if p.PathDir() != "/prefix/a/b/c" {
		t.Errorf("PathDir() = %q, want %q", p.PathDir(), "/prefix/a/b/c")
	}
}
