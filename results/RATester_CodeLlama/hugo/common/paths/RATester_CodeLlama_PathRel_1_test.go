package paths

import (
	"fmt"
	"testing"
)

func TestPathRel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		s: "a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t/u/v/w/x/y/z",
	}
	owner := &Path{
		s: "a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t/u/v/w/x/y/z",
	}
	if got := p.PathRel(owner); got != "" {
		t.Errorf("PathRel() = %v, want %v", got, "")
	}
}
