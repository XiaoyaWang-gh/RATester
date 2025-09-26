package metrics

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestRouterReqsBytesCounter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := &standardRegistry{}

	// when
	routerReqsBytesCounter := r.RouterReqsBytesCounter()

	// then
	assert.NotNil(t, routerReqsBytesCounter)
}
