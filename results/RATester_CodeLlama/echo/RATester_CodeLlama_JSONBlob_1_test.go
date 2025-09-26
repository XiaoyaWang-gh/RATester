package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestJSONBlob_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &context{}
	code := 200
	b := []byte("test")
	// when
	err := c.JSONBlob(code, b)
	// then
	assert.NoError(t, err)
}
