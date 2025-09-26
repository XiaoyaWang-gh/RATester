package paths

import (
	"fmt"
	"testing"
)

func TestSection_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		s: "a/b/c",
	}
	if p.Section() != "a" {
		t.Errorf("expected %q but got %q", "a", p.Section())
	}
}
