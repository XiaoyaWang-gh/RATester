package aggregator

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestNewProviderAggregator_1(t *testing.T) {
	type args struct {
		conf static.Providers
	}
	tests := []struct {
		name string
		args args
		want ProviderAggregator
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

			if got := NewProviderAggregator(tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProviderAggregator() = %v, want %v", got, tt.want)
			}
		})
	}
}
