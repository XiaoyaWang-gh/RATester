package nomad

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildConfig_1(t *testing.T) {
	type args struct {
		ctx   context.Context
		items []item
	}
	tests := []struct {
		name string
		args args
		want *dynamic.Configuration
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

			p := &Provider{}
			if got := p.buildConfig(tt.args.ctx, tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.buildConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
