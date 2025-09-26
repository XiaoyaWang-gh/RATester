package schema

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewDecoder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	decoder := NewDecoder()

	// Assert
	assert.NotNil(t, decoder)
}
