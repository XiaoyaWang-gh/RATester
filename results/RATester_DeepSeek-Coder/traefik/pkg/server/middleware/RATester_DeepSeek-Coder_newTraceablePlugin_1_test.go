package middleware

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/plugins"
)

func TestNewTraceablePlugin_1(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
		plug plugins.Constructor
		next http.Handler
	}
	tests := []struct {
		name    string
		args    args
		want    *traceablePlugin
		wantErr bool
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

			got, err := newTraceablePlugin(tt.args.ctx, tt.args.name, tt.args.plug, tt.args.next)
			if (err != nil) != tt.wantErr {
				t.Errorf("newTraceablePlugin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTraceablePlugin() = %v, want %v", got, tt.want)
			}
		})
	}
}
