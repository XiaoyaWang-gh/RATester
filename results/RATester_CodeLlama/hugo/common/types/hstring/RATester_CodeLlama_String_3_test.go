package hstring

import (
	"fmt"
	"testing"
)

func TestString_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := HTML("foo")
	if s.String() != "foo" {
		t.Errorf("expected %q, got %q", "foo", s.String())
	}
}
