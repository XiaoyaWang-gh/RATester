package metrics

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestTLSCertsNotAfterTimestampGauge_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := &standardRegistry{}

	// when
	result := r.TLSCertsNotAfterTimestampGauge()

	// then
	assert.NotNil(t, result)
}
