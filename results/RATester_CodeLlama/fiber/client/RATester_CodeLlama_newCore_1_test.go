package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewCore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := newCore()
	assert.NotNil(t, c)
}
