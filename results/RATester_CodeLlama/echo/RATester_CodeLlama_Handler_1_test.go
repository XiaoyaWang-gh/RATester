package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &context{}
	// when
	handler := c.Handler()
	// then
	assert.Nil(t, handler)
}
