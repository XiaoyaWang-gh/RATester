package paths

import (
	"fmt"
	"testing"
)

func TestPathNoLang_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		s: "a/b/c",
	}
	if got := p.PathNoLang(); got != "a/b/c" {
		t.Errorf("PathNoLang() = %v, want %v", got, "a/b/c")
	}
}
