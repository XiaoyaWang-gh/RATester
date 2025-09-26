package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestClone_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	l := genericResource{}

	// Act
	result := l.clone()

	// Assert
	assert.NotNil(t, result)
}
