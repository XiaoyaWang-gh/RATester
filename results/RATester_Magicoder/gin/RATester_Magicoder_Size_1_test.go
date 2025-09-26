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

	w := &responseWriter{
		size: 10,
	}

	if w.Size() != 10 {
		t.Errorf("Expected size to be 10, but got %d", w.Size())
	}
}
