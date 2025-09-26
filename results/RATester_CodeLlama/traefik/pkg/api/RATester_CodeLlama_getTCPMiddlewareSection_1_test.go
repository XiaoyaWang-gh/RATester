package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestGetTCPMiddlewareSection_1(t *testing.T) {
	type args struct {
		middlewares map[string]*runtime.TCPMiddlewareInfo
	}
	tests := []struct {
		name string
		args args
		want *section
	}{
		{
			name: "test",
			args: args{
				middlewares: map[string]*runtime.TCPMiddlewareInfo{
					"test": {
						Status: runtime.StatusDisabled,
					},
				},
			},
			want: &section{
				Total:    1,
				Warnings: 0,
				Errors:   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getTCPMiddlewareSection(tt.args.middlewares); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTCPMiddlewareSection() = %v, want %v", got, tt.want)
			}
		})
	}
}
