package media

import (
	"fmt"
	"testing"
)

func TestIsZero_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Type{}
	if !m.IsZero() {
		t.Errorf("IsZero() = %v, want %v", m.IsZero(), true)
	}
}
