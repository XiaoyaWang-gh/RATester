package hugio

import (
	"fmt"
	"testing"
)

func TestpatternLen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HasBytesWriter{
		Patterns: []*HasBytesPattern{
			{Pattern: []byte("hello")},
			{Pattern: []byte("world")},
		},
	}

	expected := 10
	actual := h.patternLen()

	if actual != expected {
		t.Errorf("Expected patternLen to return %d, but got %d", expected, actual)
	}
}
