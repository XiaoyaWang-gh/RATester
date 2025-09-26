package fiber

import (
	"fmt"
	"testing"
)

func TestRestartRouting_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	ctx := &DefaultCtx{
		app: app,
	}

	err := ctx.RestartRouting()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if ctx.indexRoute != -1 {
		t.Errorf("Expected indexRoute to be -1, got %v", ctx.indexRoute)
	}
}
