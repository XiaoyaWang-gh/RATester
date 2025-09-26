package service

import (
	"fmt"
	"testing"
)

func TestNewBufferPool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	result := newBufferPool()
	// Assert
	if result == nil {
		t.Errorf("newBufferPool() = %v, want %v", result, nil)
	}
}
