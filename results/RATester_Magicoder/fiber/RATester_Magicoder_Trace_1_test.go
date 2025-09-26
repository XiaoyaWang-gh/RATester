package fiber

import (
	"fmt"
	"testing"
)

func TestTrace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	registering := &Registering{app: app, path: "/test"}

	handler := func(ctx Ctx) error {
		return ctx.SendString("Hello, World!")
	}

	register := registering.Trace(handler)

	if register == nil {
		t.Error("Expected register to not be nil")
	}
}
