package hstrings

import (
	"fmt"
	"testing"
)

func TestString_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := StringEqualFold("foo")
	if s.String() != "foo" {
		t.Errorf("expected %q, got %q", "foo", s.String())
	}
}
