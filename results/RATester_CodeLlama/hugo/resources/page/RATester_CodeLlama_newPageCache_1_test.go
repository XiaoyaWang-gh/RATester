package page

import (
	"fmt"
	"testing"
)

func TestNewPageCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	result := newPageCache()
	// Assert
	if result == nil {
		t.Errorf("newPageCache() failed")
	}
}
