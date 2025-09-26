package paths

import (
	"fmt"
	"testing"
)

func TestContainer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		s: "a/b/c",
	}
	if p.Container() != "a/b" {
		t.Errorf("Expected %q but got %q", "a/b", p.Container())
	}
}
