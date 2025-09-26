package paths

import (
	"fmt"
	"testing"
)

func TestBaseNameNoIdentifier_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		s: "a/b/c/d/e/f",
	}
	if got := p.BaseNameNoIdentifier(); got != "f" {
		t.Errorf("BaseNameNoIdentifier() = %v, want %v", got, "f")
	}
}
