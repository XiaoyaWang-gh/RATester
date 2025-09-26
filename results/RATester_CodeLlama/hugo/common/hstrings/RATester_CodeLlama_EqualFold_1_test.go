package hstrings

import (
	"fmt"
	"testing"
)

func TestEqualFold_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := StringEqualFold("foo")
	if !s.EqualFold("foo") {
		t.Errorf("EqualFold(\"foo\") = false, want true")
	}
	if s.EqualFold("bar") {
		t.Errorf("EqualFold(\"bar\") = true, want false")
	}
}
