package metrics

import (
	"fmt"
	"reflect"
	"testing"

	"go.opentelemetry.io/otel/metric"
)

func TestNewOTLPGaugeFrom_1(t *testing.T) {
	type args struct {
		meter metric.Meter
		name  string
		desc  string
		unit  string
	}
	tests := []struct {
		name string
		args args
		want *otelGauge
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

			got := newOTLPGaugeFrom(tt.args.meter, tt.args.name, tt.args.desc, tt.args.unit)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newOTLPGaugeFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
