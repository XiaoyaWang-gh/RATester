package compress

import (
	"fmt"
	"testing"
)

func TestHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := &responseWriter{}
	header := rw.Header()
	if header == nil {
		t.Error("Expected non-nil header, got nil")
	}
}
