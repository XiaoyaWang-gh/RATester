package docshelper

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestAddDocProviderFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var fn func() DocProvider
	var docProviderFuncs []DocProviderFunc
	// Act
	AddDocProviderFunc(fn)
	// Assert
	assert.Equal(t, []DocProviderFunc{fn}, docProviderFuncs)
}
