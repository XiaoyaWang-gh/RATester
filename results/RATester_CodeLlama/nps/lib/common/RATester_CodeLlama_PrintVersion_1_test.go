package common

import (
	"fmt"
	"testing"
)

func TestPrintVersion_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	PrintVersion()
	// Assert
	// TODO: How to assert the output of PrintVersion()?
}
