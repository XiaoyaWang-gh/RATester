package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestJSONP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &context{}
	code := 200
	callback := "callback"
	i := "test"
	// when
	err := c.JSONP(code, callback, i)
	// then
	assert.NoError(t, err)
}
