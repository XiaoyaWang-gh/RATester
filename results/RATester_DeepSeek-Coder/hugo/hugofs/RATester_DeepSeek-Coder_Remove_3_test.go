package hugofs

import (
	"fmt"
	"testing"
)

func TestRemove_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := noOpFs{}
	err := fs.Remove("testfile")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
