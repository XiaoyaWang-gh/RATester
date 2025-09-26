package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSlice_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var p *pageState
	var items any

	// Act
	result, err := p.Slice(items)

	// Assert
	assert.Nil(t, result)
	assert.NotNil(t, err)
}
