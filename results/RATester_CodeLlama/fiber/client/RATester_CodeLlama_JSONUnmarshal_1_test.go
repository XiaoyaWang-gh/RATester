package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestJSONUnmarshal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &Client{}
	// when
	actual := c.JSONUnmarshal()
	// then
	assert.NotNil(t, actual)
}
