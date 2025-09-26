package paths

import (
	"fmt"
	"testing"
)

func TestBaseNoLeadingSlash_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		s: "/foo/bar",
	}
	if got := p.BaseNoLeadingSlash(); got != "bar" {
		t.Errorf("BaseNoLeadingSlash() = %v, want %v", got, "bar")
	}
}
