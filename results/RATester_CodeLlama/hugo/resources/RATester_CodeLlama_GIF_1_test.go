package resources

import (
	"fmt"
	"testing"
)

func TestGIF_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var g *giphy
	// Act
	result := g.GIF()
	// Assert
	if result != nil {
		t.Errorf("GIF() = %v, want nil", result)
	}
}
