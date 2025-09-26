package hugofs

import (
	"fmt"
	"testing"
)

func TestRemoveAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := noOpFs{}
	path := "/test"
	err := fs.RemoveAll(path)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
