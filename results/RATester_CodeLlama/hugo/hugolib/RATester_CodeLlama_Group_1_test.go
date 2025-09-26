package hugolib

import (
	"fmt"
	"testing"
)

func TestGroup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var p *pageState
	var key any
	var in any
	// Act
	_, err := p.Group(key, in)
	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
