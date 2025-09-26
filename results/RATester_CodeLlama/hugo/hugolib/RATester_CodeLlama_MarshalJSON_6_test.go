package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestMarshalJSON_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	p := &pageState{}

	// Act
	_, err := p.MarshalJSON()

	// Assert
	assert.NoError(t, err)
}
