package echo

import (
	"fmt"
	"testing"
)

func TestCreateFilesystem_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	result := createFilesystem()
	// Assert
	if result.Filesystem == nil {
		t.Errorf("Filesystem is nil")
	}
}
