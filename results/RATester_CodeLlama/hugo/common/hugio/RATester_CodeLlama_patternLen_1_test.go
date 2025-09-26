package hugio

import (
	"fmt"
	"testing"
)

func TestPatternLen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	h := &HasBytesWriter{}
	h.Patterns = []*HasBytesPattern{
		{
			Match:   true,
			Pattern: []byte("foo"),
		},
		{
			Match:   true,
			Pattern: []byte("bar"),
		},
	}

	// Act
	l := h.patternLen()

	// Assert
	if l != 6 {
		t.Errorf("Expected 6, got %d", l)
	}
}
