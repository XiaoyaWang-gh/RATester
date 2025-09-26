package pandoc

import (
	"fmt"
	"testing"
)

func TestGetPandocBinaryName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	pandocBinary := "pandoc"
	// Act
	result := getPandocBinaryName()
	// Assert
	if result != pandocBinary {
		t.Errorf("Expected %s, got %s", pandocBinary, result)
	}
}
