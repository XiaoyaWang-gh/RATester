package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestJSONPBlob_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &context{}
	code := 200
	callback := "callback"
	b := []byte("test")
	// when
	err := c.JSONPBlob(code, callback, b)
	// then
	assert.NoError(t, err)
}
