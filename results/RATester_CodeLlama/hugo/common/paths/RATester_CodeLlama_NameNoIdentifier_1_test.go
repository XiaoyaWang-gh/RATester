package paths

import (
	"fmt"
	"testing"
)

func TestNameNoIdentifier_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		s: "a/b/c/d",
	}
	if got := p.NameNoIdentifier(); got != "d" {
		t.Errorf("NameNoIdentifier() = %v, want %v", got, "d")
	}
}
