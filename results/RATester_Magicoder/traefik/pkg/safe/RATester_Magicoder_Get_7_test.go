package safe

import (
	"fmt"
	"testing"
)

func TestGet_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Safe{value: 42}
	got := s.Get()
	if got != 42 {
		t.Errorf("Safe.Get() = %v, want 42", got)
	}
}
