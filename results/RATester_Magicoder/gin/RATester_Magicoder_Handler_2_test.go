package gin

import (
	"fmt"
	"testing"
)

func TestHandler_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{
		UseH2C: true,
	}

	handler := engine.Handler()

	if handler == nil {
		t.Error("Expected a non-nil handler, but got nil")
	}

	engine.UseH2C = false
	handler = engine.Handler()

	if handler != engine {
		t.Error("Expected the handler to be the engine when UseH2C is false, but got a different handler")
	}
}
