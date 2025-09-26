package middleware

import (
	"context"
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/containous/alice"
	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/metrics"
	"github.com/traefik/traefik/v3/pkg/middlewares/accesslog"
	"github.com/traefik/traefik/v3/pkg/tracing"
)

func TestBuildEPChain_1(t *testing.T) {
	type fields struct {
		config                 static.Configuration
		accessLoggerMiddleware *accesslog.Handler
		metricsRegistry        metrics.Registry
		semConvMetricRegistry  *metrics.SemConvMetricsRegistry
		tracer                 *tracing.Tracer
		tracerCloser           io.Closer
	}
	type args struct {
		ctx            context.Context
		entryPointName string
		resourceName   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   alice.Chain
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
			if got := o.BuildEPChain(tt.args.ctx, tt.args.entryPointName, tt.args.resourceName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ObservabilityMgr.BuildEPChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
