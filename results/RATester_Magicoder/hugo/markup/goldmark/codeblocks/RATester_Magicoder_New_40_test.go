package codeblocks

import (
	"fmt"
	"testing"
)

func TestNew_40(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ext := New()
	if ext == nil {
		t.Fatal("Expected New() to return a non-nil Extender")
	}
}
