package images

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/disintegration/gift"
)

func TestInvert_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	filters := &Filters{}
	expected := gift.Invert()

	// Act
	actual := filters.Invert()

	// Assert
	assert.Equal(t, expected, actual)
}
