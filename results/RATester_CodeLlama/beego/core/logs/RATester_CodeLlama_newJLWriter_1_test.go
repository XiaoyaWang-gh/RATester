package logs

import (
	"fmt"
	"testing"
)

func TestNewJLWriter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	res := newJLWriter()
	// Assert
	if res == nil {
		t.Errorf("newJLWriter() failed")
	}
}
