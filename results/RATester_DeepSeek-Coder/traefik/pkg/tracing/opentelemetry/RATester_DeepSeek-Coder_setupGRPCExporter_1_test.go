package opentelemetry

import (
	"fmt"
	"reflect"
	"testing"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
)

func TestSetupGRPCExporter_1(t *testing.T) {
	type args struct {
		c *Config
	}
	tests := []struct {
		name    string
		args    args
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

			got, err := tt.args.c.setupGRPCExporter()
			if (err != nil) != tt.wantErr {
				t.Errorf("setupGRPCExporter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupGRPCExporter() = %v, want %v", got, tt.want)
			}
		})
	}
}
