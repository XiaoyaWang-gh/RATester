package echo

import (
	"fmt"
	"testing"
)

func TestAdd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Echo{}

	method := "GET"
	path := "/test"
	handler := func(c Context) error { return nil }

	route := e.Add(method, path, handler)

	if route == nil {
		t.Error("Expected route to be not nil")
	}

	if route.Method != method {
		t.Errorf("Expected method to be %s, but got %s", method, route.Method)
	}

	if route.Path != path {
		t.Errorf("Expected path to be %s, but got %s", path, route.Path)
	}
}
