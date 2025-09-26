package buffers

import (
	"fmt"
	"testing"
)

func TestGet_23(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	result := Get()
	// Assert
	if result == nil {
		t.Errorf("Get() = nil, want not nil")
	}
}
