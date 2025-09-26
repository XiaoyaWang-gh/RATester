package metrics

import (
	"fmt"
	"reflect"
	"testing"

	"go.opentelemetry.io/otel/metric"
)

func TestNewOTLPCounterFrom_1(t *testing.T) {
	type args struct {
		meter metric.Meter
		name  string
		desc  string
	}
	tests := []struct {
		name string
		args args
		want *otelCounter
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

			if got := newOTLPCounterFrom(tt.args.meter, tt.args.name, tt.args.desc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newOTLPCounterFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
