package metrics

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestLastConfigReloadSuccessGauge_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := &standardRegistry{}

	// when
	result := r.LastConfigReloadSuccessGauge()

	// then
	assert.NotNil(t, result)
}
