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

	w := &responseWriter{
		size: 10,
	}
	if w.Written() {
		t.Error("Written() should return false")
	}
}
