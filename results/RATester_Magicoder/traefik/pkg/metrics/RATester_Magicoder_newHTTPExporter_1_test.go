package metrics

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

func TestNewHTTPExporter_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		config *types.OtelHTTP
	}
	tests := []struct {
		name    string
		args    args
		want    sdkmetric.Exporter
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

			got, err := newHTTPExporter(tt.args.ctx, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("newHTTPExporter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newHTTPExporter() = %v, want %v", got, tt.want)
			}
		})
	}
}
