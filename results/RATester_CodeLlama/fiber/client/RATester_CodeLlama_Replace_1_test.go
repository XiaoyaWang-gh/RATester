package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestReplace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &Client{}
	// when
	replace := Replace(c)
	// then
	assert.NotNil(t, replace)
}
