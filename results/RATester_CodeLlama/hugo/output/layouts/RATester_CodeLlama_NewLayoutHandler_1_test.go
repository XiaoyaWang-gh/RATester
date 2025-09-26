package layouts

import (
	"fmt"
	"testing"
)

func TestNewLayoutHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := NewLayoutHandler()
	if handler == nil {
		t.Error("handler is nil")
	}
}
