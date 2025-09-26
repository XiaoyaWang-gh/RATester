package testhelpers

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestBackendServerUpGauge_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &CollectingHealthCheckMetrics{}
	assert.Equal(t, m.Gauge, m.BackendServerUpGauge())
}
