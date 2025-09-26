package gin

import (
	"fmt"
	"testing"
)

func TestLogger_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := Logger()
	if handler == nil {
		t.Error("Expected a HandlerFunc, but got nil")
	}
}
