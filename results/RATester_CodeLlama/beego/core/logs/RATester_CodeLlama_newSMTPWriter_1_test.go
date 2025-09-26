package logs

import (
	"fmt"
	"testing"
)

func TestNewSMTPWriter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	res := newSMTPWriter()
	// Assert
	if res == nil {
		t.Errorf("newSMTPWriter() failed")
	}
}
