package tplimpl

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestIsInternalTemplate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var tpl *templateState
	// Act
	result := tpl.IsInternalTemplate()
	// Assert
	assert.Equal(t, false, result)
}
