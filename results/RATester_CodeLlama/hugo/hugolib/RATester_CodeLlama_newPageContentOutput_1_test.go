package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewPageContentOutput_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	po := &pageOutput{}

	// Act
	cp, err := newPageContentOutput(po)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, cp)
	assert.Equal(t, po, cp.po)
}
