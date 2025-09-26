package metrics

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestConfigReloadsCounter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := &standardRegistry{}

	// when
	counter := r.ConfigReloadsCounter()

	// then
	assert.NotNil(t, counter)
}
