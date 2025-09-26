package metrics

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestServiceReqsBytesCounter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := &standardRegistry{}

	// when
	actual := r.ServiceReqsBytesCounter()

	// then
	assert.NotNil(t, actual)
}
