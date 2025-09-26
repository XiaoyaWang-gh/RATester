package paths

import (
	"fmt"
	"testing"
)

func TestIdentifierBase_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		s: "a/b/c",
	}
	if got := p.IdentifierBase(); got != "c" {
		t.Errorf("IdentifierBase() = %v, want %v", got, "c")
	}
}
