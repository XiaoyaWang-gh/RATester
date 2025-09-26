package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestR_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given: A Client
	c := &Client{}

	// when: Call R()
	r := c.R()

	// then: The returned Request should be not nil
	assert.NotNil(t, r)
}
