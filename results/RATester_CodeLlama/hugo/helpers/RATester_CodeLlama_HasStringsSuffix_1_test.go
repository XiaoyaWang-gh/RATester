package helpers

import (
	"fmt"
	"testing"
)

func TestHasStringsSuffix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := []string{"a", "b", "c"}
	suffix := []string{"c", "d"}
	if HasStringsSuffix(s, suffix) {
		t.Errorf("HasStringsSuffix(%v, %v) = true, want false", s, suffix)
	}
	s = []string{"a", "b", "c", "d"}
	suffix = []string{"c", "d"}
	if !HasStringsSuffix(s, suffix) {
		t.Errorf("HasStringsSuffix(%v, %v) = false, want true", s, suffix)
	}
	s = []string{"a", "b", "c", "d", "e"}
	suffix = []string{"c", "d"}
	if HasStringsSuffix(s, suffix) {
		t.Errorf("HasStringsSuffix(%v, %v) = true, want false", s, suffix)
	}
	s = []string{"a", "b", "c", "d", "e", "f"}
	suffix = []string{"c", "d"}
	if !HasStringsSuffix(s, suffix) {
		t.Errorf("HasStringsSuffix(%v, %v) = false, want true", s, suffix)
	}
}
