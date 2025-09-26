package echo

import (
	"fmt"
	"testing"
)

func TestRouters_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Echo{}
	e.routers = make(map[string]*Router)
	e.routers["test"] = &Router{}

	routers := e.Routers()
	if len(routers) != 1 {
		t.Errorf("Expected 1 router, got %d", len(routers))
	}

	if _, ok := routers["test"]; !ok {
		t.Errorf("Expected router with key 'test', but it was not found")
	}
}
