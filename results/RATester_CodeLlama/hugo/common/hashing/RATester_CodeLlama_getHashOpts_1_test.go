package hashing

import (
	"fmt"
	"testing"
)

func TestGetHashOpts_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	result := getHashOpts()
	// Assert
	if result == nil {
		t.Errorf("getHashOpts() = nil, want not nil")
	}
}
