package validation

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestDefaultMessage_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := Length{N: 10}
	assert.Equal(t, "Length must be exactly 10", l.DefaultMessage())
}
