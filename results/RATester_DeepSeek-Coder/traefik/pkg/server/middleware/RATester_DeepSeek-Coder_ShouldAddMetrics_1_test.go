package middleware

import (
	"fmt"
	"io"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/metrics"
	"github.com/traefik/traefik/v3/pkg/middlewares/accesslog"
	"github.com/traefik/traefik/v3/pkg/tracing"
)

func TestShouldAddMetrics_1(t *testing.T) {
	type fields struct {
		config                 static.Configuration
		accessLoggerMiddleware *accesslog.Handler
		metricsRegistry        metrics.Registry
		semConvMetricRegistry  *metrics.SemConvMetricsRegistry
		tracer                 *tracing.Tracer
		tracerCloser           io.Closer
	}
	type args struct {
		resourceName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
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

			o := &ObservabilityMgr{
				config:                 tt.fields.config,
				accessLoggerMiddleware: tt.fields.accessLoggerMiddleware,
				metricsRegistry:        tt.fields.metricsRegistry,
				semConvMetricRegistry:  tt.fields.semConvMetricRegistry,
				tracer:                 tt.fields.tracer,
				tracerCloser:           tt.fields.tracerCloser,
			}
			if got := o.ShouldAddMetrics(tt.args.resourceName); got != tt.want {
				t.Errorf("ObservabilityMgr.ShouldAddMetrics() = %v, want %v", got, tt.want)
			}
		})
	}
}
