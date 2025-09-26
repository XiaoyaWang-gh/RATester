package hugolib

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestParseError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	p := &pageState{}
	err := errors.New("error")
	input := []byte("input")
	offset := 1

	// Act
	actual := p.parseError(err, input, offset)

	// Assert
	assert.Equal(t, err, actual)
}
