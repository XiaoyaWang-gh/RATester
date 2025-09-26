package fiber

import (
	"fmt"
	"testing"
)

func TestMethod_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var c *DefaultCtx
	var override []string
	var method string

	// TEST CASE 1:
	// Method()
	method = c.Method()
	if method != "" {
		t.Errorf("Method() = %v, want %v", method, "")
	}

	// TEST CASE 2:
	// Method(override...)
	override = []string{"GET"}
	method = c.Method(override...)
	if method != "GET" {
		t.Errorf("Method(override...) = %v, want %v", method, "GET")
	}
}
