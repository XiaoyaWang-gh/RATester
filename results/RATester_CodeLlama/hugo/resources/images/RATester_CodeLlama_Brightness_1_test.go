package images

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/disintegration/gift"
)

func TestBrightness_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var percentage any = 100
	var filters Filters
	var filter gift.Filter

	// Act
	filter = filters.Brightness(percentage)

	// Assert
	assert.NotNil(t, filter)
}
