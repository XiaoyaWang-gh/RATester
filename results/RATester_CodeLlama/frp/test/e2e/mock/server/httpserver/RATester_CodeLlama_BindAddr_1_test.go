package httpserver

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestBindAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Server{
		bindAddr: "127.0.0.1",
	}
	assert.Equal(t, "127.0.0.1", s.BindAddr())
}
