package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	l := &genericResource{}

	// Act
	actual := l.size()

	// Assert
	assert.Equal(t, int64(0), actual)
}
