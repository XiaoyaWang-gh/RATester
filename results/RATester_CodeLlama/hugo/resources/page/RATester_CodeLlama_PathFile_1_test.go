package page

import (
	"fmt"
	"testing"
)

func TestPathFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pagePathBuilder{}
	p.Add("a", "b", "c")
	p.ConcatLast("d")
	p.Sanitize()
	if p.PathFile() != "a/b/c/d" {
		t.Errorf("p.PathFile() = %q, want %q", p.PathFile(), "a/b/c/d")
	}
}
