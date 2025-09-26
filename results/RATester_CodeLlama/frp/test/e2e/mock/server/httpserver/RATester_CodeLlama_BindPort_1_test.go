package httpserver

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestBindPort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Server{
		bindPort: 1234,
	}
	assert.Equal(t, 1234, s.BindPort())
}
