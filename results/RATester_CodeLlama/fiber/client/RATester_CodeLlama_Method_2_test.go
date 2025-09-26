package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestMethod_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	r.method = "GET"
	assert.Equal(t, "GET", r.Method())
}
