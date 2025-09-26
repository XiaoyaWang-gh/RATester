package paths

import (
	"fmt"
	"testing"
)

func TestComponent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		component: "test",
	}
	if p.Component() != "test" {
		t.Errorf("expected %q, got %q", "test", p.Component())
	}
}
