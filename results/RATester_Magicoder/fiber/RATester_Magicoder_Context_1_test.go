package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{}

	// Test case 1: Test the Context method
	expected := &fasthttp.RequestCtx{}
	ctx.fasthttp = expected
	result := ctx.Context()
	if result != expected {
		t.Errorf("Test case 1 failed: Expected %v, but got %v", expected, result)
	}

	// Test case 2: Test the Context method with nil fasthttp
	ctx.fasthttp = nil
	result = ctx.Context()
	if result != nil {
		t.Errorf("Test case 2 failed: Expected nil, but got %v", result)
	}
}
