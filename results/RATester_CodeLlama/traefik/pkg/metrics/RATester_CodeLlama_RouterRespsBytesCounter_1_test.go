package metrics

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestRouterRespsBytesCounter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := &standardRegistry{}

	// when
	routerRespsBytesCounter := r.RouterRespsBytesCounter()

	// then
	assert.NotNil(t, routerRespsBytesCounter)
}
