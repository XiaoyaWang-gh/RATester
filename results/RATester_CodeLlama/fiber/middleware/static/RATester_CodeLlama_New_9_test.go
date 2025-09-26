package static

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNew_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	root := "."
	cfg := []Config{}
	// Act
	handler := New(root, cfg...)
	// Assert
	assert.NotNil(t, handler)
}
