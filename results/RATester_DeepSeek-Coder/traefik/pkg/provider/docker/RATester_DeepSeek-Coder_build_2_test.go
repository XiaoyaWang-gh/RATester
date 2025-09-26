package docker

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuild_2(t *testing.T) {
	type args struct {
		ctx                 context.Context
		containersInspected []dockerData
	}
	tests := []struct {
		name string
		p    *DynConfBuilder
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

			if got := tt.p.build(tt.args.ctx, tt.args.containersInspected); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DynConfBuilder.build() = %v, want %v", got, tt.want)
			}
		})
	}
}
