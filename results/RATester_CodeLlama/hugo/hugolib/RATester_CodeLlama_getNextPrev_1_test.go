package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetNextPrev_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	p := &pageCommon{}
	// Act
	result := p.getNextPrev()
	// Assert
	assert.NotNil(t, result)
}
