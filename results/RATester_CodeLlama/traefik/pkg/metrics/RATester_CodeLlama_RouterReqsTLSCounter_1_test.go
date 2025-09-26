package metrics

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestRouterReqsTLSCounter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := &standardRegistry{}

	// when
	routerReqsTLSCounter := r.RouterReqsTLSCounter()

	// then
	assert.NotNil(t, routerReqsTLSCounter)
}
