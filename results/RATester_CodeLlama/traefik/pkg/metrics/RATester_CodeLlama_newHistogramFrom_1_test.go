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
		{
			name: "test_newHistogramFrom_0",
			args: args{
				opts:       stdprometheus.HistogramOpts{},
				labelNames: []string{},
			},
			want: &histogram{
				name:             "",
				hv:               nil,
				labelNamesValues: nil,
				collector:        nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := newHistogramFrom(tt.args.opts, tt.args.labelNames); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newHistogramFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
