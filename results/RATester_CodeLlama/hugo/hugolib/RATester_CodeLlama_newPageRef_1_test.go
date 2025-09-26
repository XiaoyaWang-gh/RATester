package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewPageRef_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var p *pageState

	// Act
	actual := newPageRef(p)

	// Assert
	assert.Equal(t, pageRef{p: p}, actual)
}
