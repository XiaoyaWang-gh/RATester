package metrics

import (
	"fmt"
	"reflect"
	"testing"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func TestNewCounterWithHeadersFrom_1(t *testing.T) {
	type args struct {
		opts       stdprometheus.CounterOpts
		headers    map[string]string
		labelNames []string
	}
	tests := []struct {
		name string
		args args
		want *counterWithHeaders
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

			if got := newCounterWithHeadersFrom(tt.args.opts, tt.args.headers, tt.args.labelNames); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCounterWithHeadersFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
