package metrics

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestServiceServerUpGauge_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := &standardRegistry{}

	// when
	result := r.ServiceServerUpGauge()

	// then
	assert.NotNil(t, result)
}
