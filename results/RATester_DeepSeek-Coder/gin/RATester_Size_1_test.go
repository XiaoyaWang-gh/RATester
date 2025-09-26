package gin

import (
	"fmt"
	"testing"
)

func TestSize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &responseWriter{size: 100}
	if w.Size() != 100 {
		t.Errorf("Expected size to be 100, got %d", w.Size())
	}
}
