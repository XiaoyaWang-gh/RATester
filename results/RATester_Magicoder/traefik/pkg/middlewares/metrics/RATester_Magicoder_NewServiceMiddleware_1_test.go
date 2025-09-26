package metrics

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/metrics"
)

func TestNewServiceMiddleware_1(t *testing.T) {
	type args struct {
		ctx         context.Context
		next        http.Handler
		registry    metrics.Registry
		serviceName string
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

			if got := NewServiceMiddleware(tt.args.ctx, tt.args.next, tt.args.registry, tt.args.serviceName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServiceMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
