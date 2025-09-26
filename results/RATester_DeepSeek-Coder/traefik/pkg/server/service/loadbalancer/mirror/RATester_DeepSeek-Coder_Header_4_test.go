package mirror

import (
	"fmt"
	"testing"
)

func TestHeader_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := blackHoleResponseWriter{}
	header := b.Header()
	if header == nil {
		t.Errorf("Expected non-nil header, got nil")
	}
}
