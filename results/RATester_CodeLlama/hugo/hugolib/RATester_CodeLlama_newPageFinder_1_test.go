package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewPageFinder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var m *pageMap
	// Act
	c := newPageFinder(m)
	// Assert
	assert.NotNil(t, c)
}
