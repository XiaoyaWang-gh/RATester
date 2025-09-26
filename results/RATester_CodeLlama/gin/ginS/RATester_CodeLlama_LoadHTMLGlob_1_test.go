package ginS

import (
	"fmt"
	"testing"
)

func TestLoadHTMLGlob_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	// Arrange
	pattern := "testdata/*.html"

	// Act
	LoadHTMLGlob(pattern)

	// Assert
	// TODO
}
