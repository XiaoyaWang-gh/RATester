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

	ctx := &DefaultCtx{
		method: "GET",
	}

	result := ctx.Method()
	if result != "GET" {
		t.Errorf("Expected 'GET', got '%s'", result)
	}

	result = ctx.Method("POST")
	if result != "POST" {
		t.Errorf("Expected 'POST', got '%s'", result)
	}
}
