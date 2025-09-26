package metrics

import (
	"fmt"
	"reflect"
	"testing"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func TestNewHistogramFrom_1(t *testing.T) {
	type args struct {
		opts       stdprometheus.HistogramOpts
		labelNames []string
	}
	tests := []struct {
		name string
		args args
		want *histogram
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

			got := newHistogramFrom(tt.args.opts, tt.args.labelNames)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newHistogramFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
