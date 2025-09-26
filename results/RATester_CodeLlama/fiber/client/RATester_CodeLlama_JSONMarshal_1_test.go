package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestJSONMarshal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &Client{}
	// when
	result := c.JSONMarshal()
	// then
	assert.NotNil(t, result)
}
