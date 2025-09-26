package metrics

import (
	"fmt"
	"testing"
)

func TestNewDynamicConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	result := newDynamicConfig()
	// Assert
	if result == nil {
		t.Errorf("newDynamicConfig() failed")
	}
}
