package idempotency

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewMemoryLock_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	result := NewMemoryLock()
	// Assert
	assert.NotNil(t, result)
}
