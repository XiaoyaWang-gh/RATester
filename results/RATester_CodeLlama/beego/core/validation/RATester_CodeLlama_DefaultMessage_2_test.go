package validation

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestDefaultMessage_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := MaxSize{Max: 10}
	assert.Equal(t, "MaxSize must be at most 10", m.DefaultMessage())
}
