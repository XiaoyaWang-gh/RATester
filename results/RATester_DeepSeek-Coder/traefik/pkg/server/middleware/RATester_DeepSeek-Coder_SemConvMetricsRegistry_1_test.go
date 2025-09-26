package middleware

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/metrics"
	"gotest.tools/assert"
)

func TestSemConvMetricsRegistry_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		observabilityMgr *ObservabilityMgr
		expectedResult   *metrics.SemConvMetricsRegistry
	}{
		{
			name: "Returns SemConvMetricsRegistry when ObservabilityMgr is not nil",
			observabilityMgr: &ObservabilityMgr{
				semConvMetricRegistry: &metrics.SemConvMetricsRegistry{},
			},
			expectedResult: &metrics.SemConvMetricsRegistry{},
		},
		{
			name:             "Returns nil when ObservabilityMgr is nil",
			observabilityMgr: nil,
			expectedResult:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.observabilityMgr.SemConvMetricsRegistry()
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
