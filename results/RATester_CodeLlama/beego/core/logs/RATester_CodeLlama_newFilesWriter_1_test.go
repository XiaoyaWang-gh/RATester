package logs

import (
	"fmt"
	"testing"
)

func TestNewFilesWriter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	res := newFilesWriter()
	// Assert
	if res == nil {
		t.Errorf("newFilesWriter() failed")
	}
}
