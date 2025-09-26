package common

import (
	"fmt"
	"testing"
)

func TestNewPool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	// Assert
}
