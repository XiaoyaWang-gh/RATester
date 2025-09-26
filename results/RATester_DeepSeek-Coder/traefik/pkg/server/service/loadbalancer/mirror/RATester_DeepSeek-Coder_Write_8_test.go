package mirror

import (
	"fmt"
	"testing"
)

func TestWrite_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := blackHoleResponseWriter{}
	data := []byte("test data")
	n, err := b.Write(data)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if n != len(data) {
		t.Errorf("Expected to write %d bytes, wrote %d", len(data), n)
	}
}
