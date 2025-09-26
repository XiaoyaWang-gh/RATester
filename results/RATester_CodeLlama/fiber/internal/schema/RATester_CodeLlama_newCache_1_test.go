package schema

import (
	"fmt"
	"testing"
)

func TestNewCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	c := newCache()
	// Assert
	if c == nil {
		t.Errorf("newCache() = %v, want %v", c, nil)
	}
}
