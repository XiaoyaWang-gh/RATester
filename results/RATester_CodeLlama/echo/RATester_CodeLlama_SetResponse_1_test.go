package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestSetResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &context{}
	r := &Response{}

	// when
	c.SetResponse(r)

	// then
	assert.Equal(t, r, c.response)
}
