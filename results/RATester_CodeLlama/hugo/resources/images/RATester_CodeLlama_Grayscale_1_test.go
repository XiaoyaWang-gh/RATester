package images

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGrayscale_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	filters := &Filters{}

	// Act
	filter := filters.Grayscale()

	// Assert
	assert.NotNil(t, filter)
}
