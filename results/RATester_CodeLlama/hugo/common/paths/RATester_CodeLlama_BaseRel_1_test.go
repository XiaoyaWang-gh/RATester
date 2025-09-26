package paths

import (
	"fmt"
	"testing"
)

func TestBaseRel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		s: "/a/b/c",
	}
	owner := &Path{
		s: "/a/b",
	}
	if got := p.BaseRel(owner); got != "c" {
		t.Errorf("BaseRel() = %v, want %v", got, "c")
	}
}
