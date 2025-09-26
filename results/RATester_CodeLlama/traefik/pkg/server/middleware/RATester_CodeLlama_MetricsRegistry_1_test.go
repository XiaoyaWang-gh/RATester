package middleware

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestMetricsRegistry_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ObservabilityMgr{}
	assert.Equal(t, nil, o.MetricsRegistry())
}
