package metrics

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestNewSemConvMetricRegistry_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		config *types.OTLP
	}
	tests := []struct {
		name    string
		args    args
		want    *SemConvMetricsRegistry
		wantErr bool
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

			got, err := NewSemConvMetricRegistry(tt.args.ctx, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSemConvMetricRegistry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSemConvMetricRegistry() = %v, want %v", got, tt.want)
			}
		})
	}
}
