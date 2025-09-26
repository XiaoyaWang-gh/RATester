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

	e := &Echo{}
	router := NewRouter(e)

	if router.tree == nil {
		t.Errorf("Expected tree to be initialized")
	}

	if router.routes == nil {
		t.Errorf("Expected routes to be initialized")
	}

	if router.echo != e {
		t.Errorf("Expected echo to be set")
	}
}
