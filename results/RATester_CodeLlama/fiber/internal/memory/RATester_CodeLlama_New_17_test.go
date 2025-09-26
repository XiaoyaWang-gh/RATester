package memory

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNew_17(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	store := New()
	// Assert
	assert.NotNil(t, store)
}
