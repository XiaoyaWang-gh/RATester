package hugolib

import (
	"fmt"
	"testing"
)

func TestIsZero_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &cachedContent{}
	if c.IsZero() {
		t.Error("expected false")
	}
}
