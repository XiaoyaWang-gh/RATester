package middleware

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/metrics"
)

func TestSemConvMetricsRegistry_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ObservabilityMgr{
		semConvMetricRegistry: &metrics.SemConvMetricsRegistry{},
	}

	result := o.SemConvMetricsRegistry()

	if result != o.semConvMetricRegistry {
		t.Errorf("Expected %v, got %v", o.semConvMetricRegistry, result)
	}
}
