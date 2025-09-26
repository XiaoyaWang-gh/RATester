package metrics

import (
	"fmt"
	"reflect"
	"testing"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func TestNewGaugeFrom_1(t *testing.T) {
	type args struct {
		opts       stdprometheus.GaugeOpts
		labelNames []string
	}
	tests := []struct {
		name string
		args args
		want *gauge
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

			if got := newGaugeFrom(tt.args.opts, tt.args.labelNames); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newGaugeFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
