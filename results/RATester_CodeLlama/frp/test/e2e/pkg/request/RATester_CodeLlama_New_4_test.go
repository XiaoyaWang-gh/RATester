package request

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestNew_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := New()
	assert.Equal(t, "tcp", req.protocol)
	assert.Equal(t, "127.0.0.1", req.addr)
}
