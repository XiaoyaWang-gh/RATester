package pagemeta

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestIsResourceValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	s := Source{}

	// Act
	actual := s.IsResourceValue()

	// Assert
	assert.False(t, actual)
}
