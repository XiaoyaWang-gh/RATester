package pageparser

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestIsShortcodeParam_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var i Item
	// Act
	var actual bool
	// Assert
	actual = i.IsShortcodeParam()
	assert.Equal(t, false, actual)
}
