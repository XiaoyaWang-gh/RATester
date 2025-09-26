package page

import (
	"fmt"
	"testing"
)

func TestLink_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pagePathBuilder{
		els: []string{"a", "b", "c"},
		d: TargetPathDescriptor{
			BaseName: "c",
		},
	}

	if got := p.Link(); got != "/a/b/c" {
		t.Errorf("Link() = %v, want %v", got, "/a/b/c")
	}
}
