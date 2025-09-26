package udp

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/udp"
)

func TestBuildHandlers_3(t *testing.T) {
	type args struct {
		rootCtx     context.Context
		entryPoints []string
	}
	tests := []struct {
		name string
		args args
		want map[string]udp.Handler
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

			m := &Manager{}
			if got := m.BuildHandlers(tt.args.rootCtx, tt.args.entryPoints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}
