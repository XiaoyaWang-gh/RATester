package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestInline_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &context{}
	file := "test.txt"
	name := "test.txt"
	// when
	err := c.Inline(file, name)
	// then
	assert.NoError(t, err)
}
