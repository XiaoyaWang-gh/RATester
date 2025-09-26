package ginS

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestStatic_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	// Arrange
	relativePath := "test"
	root := "test"
	expected := engine().Static(relativePath, root)

	// Act
	actual := Static(relativePath, root)

	// Assert
	assert.Equal(t, expected, actual)
}
