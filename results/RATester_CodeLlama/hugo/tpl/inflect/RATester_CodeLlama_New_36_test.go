package inflect

import (
	"fmt"
	"testing"
)

func TestNew_36(t *testing.T) {
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
