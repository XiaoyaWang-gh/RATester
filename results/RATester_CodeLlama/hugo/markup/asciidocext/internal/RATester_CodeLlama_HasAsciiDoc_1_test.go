package internal

import (
	"fmt"
	"testing"
)

func TestHasAsciiDoc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if !HasAsciiDoc() {
		t.Errorf("Expected %s to be in path", asciiDocBinaryName)
	}
}
