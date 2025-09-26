package echo

import (
	"fmt"
	"testing"
)

func TestNewRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	router := NewRouter(e)

	if router.tree == nil {
		t.Error("Expected tree to be initialized")
	}

	if router.routes == nil {
		t.Error("Expected routes to be initialized")
	}

	if router.echo != e {
		t.Error("Expected echo to be set")
	}
}
