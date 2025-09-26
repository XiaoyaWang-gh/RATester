package validation

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestDefaultMessage_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Min{
		Min: 10,
		Key: "age",
	}
	assert.Equal(t, "age must be greater than or equal to 10", m.DefaultMessage())
}
