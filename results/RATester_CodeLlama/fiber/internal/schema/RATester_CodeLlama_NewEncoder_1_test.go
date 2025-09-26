package schema

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewEncoder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	encoder := NewEncoder()

	// Assert
	assert.NotNil(t, encoder)
}
