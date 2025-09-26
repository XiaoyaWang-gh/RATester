package echo

import (
	"fmt"
	"testing"
)

func TestFindRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Echo{
		routers: make(map[string]*Router),
		router:  &Router{},
	}

	// Test with an existing host
	e.routers["example.com"] = &Router{}
	r := e.findRouter("example.com")
	if r != e.routers["example.com"] {
		t.Errorf("Expected %v, got %v", e.routers["example.com"], r)
	}

	// Test with a non-existing host
	r = e.findRouter("non-existing.com")
	if r != e.router {
		t.Errorf("Expected %v, got %v", e.router, r)
	}

	// Test with an empty host
	r = e.findRouter("")
	if r != e.router {
		t.Errorf("Expected %v, got %v", e.router, r)
	}
}
