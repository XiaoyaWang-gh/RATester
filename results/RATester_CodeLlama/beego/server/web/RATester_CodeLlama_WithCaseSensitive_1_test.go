package web

import (
	"fmt"
	"testing"
)

func TestWithCaseSensitive_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	sensitive := true
	// Act
	opt := WithCaseSensitive(sensitive)
	// Assert
	if opt == nil {
		t.Errorf("WithCaseSensitive() = %v, want %v", opt, nil)
	}
}
