package blockquotes

import (
	"fmt"
	"testing"
)

func TestNew_52(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	result := New()
	// Assert
	if result == nil {
		t.Errorf("New() = %v, want %v", result, "not nil")
	}
}
