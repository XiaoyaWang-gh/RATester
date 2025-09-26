package tcp

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	router, err := NewRouter()
	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, router)
}
