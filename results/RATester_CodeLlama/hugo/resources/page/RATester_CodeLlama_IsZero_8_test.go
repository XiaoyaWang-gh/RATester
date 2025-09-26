package page

import (
	"fmt"
	"testing"
)

func TestIsZero_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := Summary{}
	if !s.IsZero() {
		t.Errorf("IsZero() = %v, want %v", s.IsZero(), true)
	}
}
