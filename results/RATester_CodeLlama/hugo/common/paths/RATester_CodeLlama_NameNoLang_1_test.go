package paths

import (
	"fmt"
	"testing"
)

func TestNameNoLang_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		s: "foo/bar/baz.md",
	}

	if got := p.NameNoLang(); got != "baz" {
		t.Errorf("NameNoLang() = %v, want %v", got, "baz")
	}
}
