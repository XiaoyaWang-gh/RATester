package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestData_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	l := &genericResource{}

	// Act
	actual := l.Data()

	// Assert
	assert.Equal(t, l.sd.Data, actual)
}
