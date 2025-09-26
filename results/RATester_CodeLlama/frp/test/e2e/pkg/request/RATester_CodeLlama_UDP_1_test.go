package request

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestUDP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	r.UDP()
	assert.Equal(t, "udp", r.protocol)
}
