package navigation

import (
	"fmt"
	"testing"
)

func TestNewMenuCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	result := newMenuCache()
	// Assert
	if result == nil {
		t.Errorf("newMenuCache() failed")
	}
}
