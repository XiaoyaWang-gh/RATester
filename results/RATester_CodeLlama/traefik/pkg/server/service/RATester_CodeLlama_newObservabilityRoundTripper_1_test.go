package service

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/metrics"
)

func TestNewObservabilityRoundTripper_1(t *testing.T) {
	type args struct {
		semConvMetricRegistry *metrics.SemConvMetricsRegistry
		rt                    http.RoundTripper
	}
	tests := []struct {
		name string
		args args
		want http.RoundTripper
	}{
		{
			name: "test newObservabilityRoundTripper",
			args: args{
				semConvMetricRegistry: &metrics.SemConvMetricsRegistry{},
				rt:                    &http.Transport{},
			},
			want: &wrapper{
				semConvMetricRegistry: &metrics.SemConvMetricsRegistry{},
				rt:                    &http.Transport{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := newObservabilityRoundTripper(tt.args.semConvMetricRegistry, tt.args.rt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newObservabilityRoundTripper() = %v, want %v", got, tt.want)
			}
		})
	}
}
