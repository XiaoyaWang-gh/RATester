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

	s := &Safe{value: "test value"}
	got := s.Get()
	want := "test value"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
