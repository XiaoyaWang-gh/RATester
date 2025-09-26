package gin

import (
	"fmt"
	"testing"
)

func TestWritten_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &responseWriter{size: 0}
	if w.Written() {
		t.Error("Expected false, got true")
	}

	w.size = 1
	if !w.Written() {
		t.Error("Expected true, got false")
	}
}
