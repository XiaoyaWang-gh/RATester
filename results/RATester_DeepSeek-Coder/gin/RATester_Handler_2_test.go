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
		t.Errorf("Expected a non-nil handler")
	}

	engine.UseH2C = false
	handler = engine.Handler()

	if handler == nil {
		t.Errorf("Expected a non-nil handler")
	}
}
