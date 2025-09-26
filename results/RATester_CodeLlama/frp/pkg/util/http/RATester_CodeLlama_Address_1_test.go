package http

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestAddress_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Server{
		addr: "127.0.0.1:8080",
	}
	assert.Equal(t, "127.0.0.1:8080", s.Address())
}
