package metrics

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestRegisterStatsd_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		config *types.Statsd
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

			if got := RegisterStatsd(tt.args.ctx, tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterStatsd() = %v, want %v", got, tt.want)
			}
		})
	}
}
