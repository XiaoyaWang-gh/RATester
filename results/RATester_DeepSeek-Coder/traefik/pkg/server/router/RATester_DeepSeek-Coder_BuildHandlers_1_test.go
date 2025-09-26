package router

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBuildHandlers_1(t *testing.T) {
	type args struct {
		rootCtx     context.Context
		entryPoints []string
		tls         bool
	}
	tests := []struct {
		name string
		args args
		want map[string]http.Handler
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
			if got := m.BuildHandlers(tt.args.rootCtx, tt.args.entryPoints, tt.args.tls); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.BuildHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}
