package paths

import (
	"fmt"
	"testing"
)

func TestTrimLeadingSlash_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := Path{s: "/foo"}
	p.TrimLeadingSlash()
	if p.s != "foo" {
		t.Errorf("expected %q but got %q", "foo", p.s)
	}
}
