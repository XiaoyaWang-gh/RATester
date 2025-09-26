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

	h := &HasBytesWriter{
		Patterns: []*HasBytesPattern{
			{Pattern: []byte("test1")},
			{Pattern: []byte("test2")},
			{Pattern: []byte("test3")},
		},
	}

	expected := len("test1") + len("test2") + len("test3")
	result := h.patternLen()

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
