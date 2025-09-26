package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestResponseHook_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &Client{}
	// when
	result := c.ResponseHook()
	// then
	assert.Equal(t, []ResponseHook{}, result)
}
