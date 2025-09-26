package opentelemetry

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
)

func TestSetupHTTPExporter_1(t *testing.T) {
	type fields struct {
		HTTP *types.OtelHTTP
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
				HTTP: tt.fields.HTTP,
			}
			got, err := c.setupHTTPExporter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.setupHTTPExporter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.setupHTTPExporter() = %v, want %v", got, tt.want)
			}
		})
	}
}
