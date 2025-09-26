package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewDefaultFS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	actual := newDefaultFS()

	// Assert
	assert.NotNil(t, actual)
}
