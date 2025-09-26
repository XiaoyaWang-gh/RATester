package resource

import (
	"fmt"
	"testing"
)

func TestMergeByLanguageInterface_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var r Resources
	var in any
	// Act
	_, err := r.MergeByLanguageInterface(in)
	// Assert
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
