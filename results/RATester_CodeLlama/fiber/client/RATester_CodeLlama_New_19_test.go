package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNew_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	client := New()
	// Assert
	assert.NotNil(t, client)
}
