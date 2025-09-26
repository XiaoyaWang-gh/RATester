package tcp

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestBuildHandlers_2(t *testing.T) {
	type args struct {
		rootCtx     context.Context
		entryPoints []string
	}
	tests := []struct {
		name string
		args args
		want map[string]*Router
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
				t.Errorf("Manager.BuildHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}
