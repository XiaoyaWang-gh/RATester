package echo

import (
	"fmt"
	"testing"
)

func TestGroup_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	g := e.Group("/group")

	if g.prefix != "/group" {
		t.Errorf("Expected prefix to be '/group', but got '%s'", g.prefix)
	}

	if g.echo != e {
		t.Errorf("Expected echo to be the same instance, but got a different instance")
	}

	if len(g.middleware) != 0 {
		t.Errorf("Expected middleware to be empty, but got %d middleware", len(g.middleware))
	}
}
