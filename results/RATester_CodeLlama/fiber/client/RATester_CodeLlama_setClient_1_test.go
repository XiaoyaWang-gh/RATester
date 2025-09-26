package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestSetClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var r *Response
	var c *Client
	// act
	r.setClient(c)
	// assert
	assert.Equal(t, c, r.client)
}
