package opentelemetry

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
)

func TestSetupGRPCExporter_1(t *testing.T) {
	type fields struct {
		GRPC *types.OtelGRPC
	}
	tests := []struct {
		name    string
		fields  fields
		want    *otlptrace.Exporter
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

			c := &Config{
				GRPC: tt.fields.GRPC,
			}
			got, err := c.setupGRPCExporter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.setupGRPCExporter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.setupGRPCExporter() = %v, want %v", got, tt.want)
			}
		})
	}
}
