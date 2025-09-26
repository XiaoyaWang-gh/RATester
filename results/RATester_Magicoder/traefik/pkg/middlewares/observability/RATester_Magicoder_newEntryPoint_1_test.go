package observability

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/metrics"
	"github.com/traefik/traefik/v3/pkg/tracing"
)

func TestNewEntryPoint_1(t *testing.T) {
	type args struct {
		ctx                   context.Context
		tracer                *tracing.Tracer
		semConvMetricRegistry *metrics.SemConvMetricsRegistry
		entryPointName        string
		next                  http.Handler
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := newEntryPoint(tt.args.ctx, tt.args.tracer, tt.args.semConvMetricRegistry, tt.args.entryPointName, tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newEntryPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
