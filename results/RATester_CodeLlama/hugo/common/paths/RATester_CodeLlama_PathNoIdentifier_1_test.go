package paths

import (
	"fmt"
	"testing"
)

func TestPathNoIdentifier_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		s: "a/b/c",
	}
	if got := p.PathNoIdentifier(); got != "a/b/c" {
		t.Errorf("PathNoIdentifier() = %v, want %v", got, "a/b/c")
	}
}
