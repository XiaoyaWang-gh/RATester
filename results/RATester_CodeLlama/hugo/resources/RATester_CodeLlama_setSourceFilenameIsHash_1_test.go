package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSetSourceFilenameIsHash_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	l := &genericResource{}
	b := true

	// Act
	l.setSourceFilenameIsHash(b)

	// Assert
	assert.Equal(t, b, l.sourceFilenameIsHash)
}
