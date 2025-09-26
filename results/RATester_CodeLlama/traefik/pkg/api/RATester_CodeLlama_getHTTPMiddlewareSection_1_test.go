package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestGetHTTPMiddlewareSection_1(t *testing.T) {
	type args struct {
		middlewares map[string]*runtime.MiddlewareInfo
	}
	tests := []struct {
		name string
		args args
		want *section
	}{
		{
			name: "test",
			args: args{
				middlewares: map[string]*runtime.MiddlewareInfo{
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

			if got := getHTTPMiddlewareSection(tt.args.middlewares); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHTTPMiddlewareSection() = %v, want %v", got, tt.want)
			}
		})
	}
}
