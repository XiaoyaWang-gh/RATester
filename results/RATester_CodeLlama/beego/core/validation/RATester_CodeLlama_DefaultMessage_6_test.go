package validation

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestDefaultMessage_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Max{Max: 10, Key: "age"}
	assert.Equal(t, "age must be less than or equal to 10", m.DefaultMessage())
}
