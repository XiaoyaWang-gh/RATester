package validation

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestDefaultMessage_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Range{
		Min: Min{
			Min: 1,
		},
		Max: Max{
			Max: 10,
		},
	}
	assert.Equal(t, "Range: 1, 10", r.DefaultMessage())
}
