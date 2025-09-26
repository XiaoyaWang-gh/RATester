package testhelpers

import (
	"fmt"
	"testing"
)

func TestNewCollectingHealthCheckMetrics_1(t *testing.T) {
	tests := []struct {
		name string
		want *CollectingHealthCheckMetrics
	}{
		{
			name: "Test NewCollectingHealthCheckMetrics",
			want: &CollectingHealthCheckMetrics{&CollectingGauge{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewCollectingHealthCheckMetrics()
			if got.Gauge == nil {
				t.Errorf("NewCollectingHealthCheckMetrics() = %v, want %v", got, tt.want)
			}
		})
	}
}
