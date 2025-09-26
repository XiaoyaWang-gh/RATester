package helpers

import (
	"fmt"
	"testing"
)

func TestHasStringsPrefix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := []string{"a", "b", "c"}
	prefix := []string{"a", "b"}
	if !HasStringsPrefix(s, prefix) {
		t.Errorf("HasStringsPrefix(%v, %v) = false, want true", s, prefix)
	}
}
