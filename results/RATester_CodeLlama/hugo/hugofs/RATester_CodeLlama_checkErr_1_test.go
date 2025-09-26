package hugofs

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestCheckErr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	w := &Walkway{}
	filename := "filename"
	err := errors.New("error")

	// When
	result := w.checkErr(filename, err)

	// Then
	assert.False(t, result)
}
