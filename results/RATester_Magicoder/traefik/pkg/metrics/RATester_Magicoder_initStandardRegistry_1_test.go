package metrics

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestInitStandardRegistry_1(t *testing.T) {
	type args struct {
		config *types.Prometheus
	}
	tests := []struct {
		name string
		args args
		want Registry
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

			if got := initStandardRegistry(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initStandardRegistry() = %v, want %v", got, tt.want)
			}
		})
	}
}
