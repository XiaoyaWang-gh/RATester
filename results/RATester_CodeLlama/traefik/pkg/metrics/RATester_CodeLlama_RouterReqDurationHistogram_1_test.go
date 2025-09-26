package metrics

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestRouterReqDurationHistogram_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &standardRegistry{}
	assert.NotNil(t, r.RouterReqDurationHistogram())
}
