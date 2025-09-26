package hugolib

import (
	"fmt"
	"testing"
)

func TestReadSourceAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var c contentParseInfo
	// Act
	result, err := c.readSourceAll()
	// Assert
	if err != nil {
		t.Errorf("readSourceAll() error = %v", err)
		return
	}
	if len(result) != 0 {
		t.Errorf("readSourceAll() result = %v, want %v", result, 0)
	}
}
