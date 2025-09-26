package metrics

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/containous/alice"
	"github.com/traefik/traefik/v3/pkg/metrics"
)

func TestWrapServiceHandler_1(t *testing.T) {
	type args struct {
		ctx         context.Context
		registry    metrics.Registry
		serviceName string
	}
	tests := []struct {
		name string
		args args
		want alice.Constructor
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

			got := WrapServiceHandler(tt.args.ctx, tt.args.registry, tt.args.serviceName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapServiceHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
