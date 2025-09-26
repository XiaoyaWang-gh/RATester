package gin

import (
	"fmt"
	"testing"
)

func TestErrorLogger_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := ErrorLogger()
	if handler == nil {
		t.Error("Expected a non-nil handler, but got nil")
	}
}
