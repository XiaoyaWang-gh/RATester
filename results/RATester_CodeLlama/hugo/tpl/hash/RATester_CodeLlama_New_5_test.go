package hash

import (
	"fmt"
	"testing"
)

func TestNew_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := New()
	if n == nil {
		t.Error("New() returned nil")
	}
}
