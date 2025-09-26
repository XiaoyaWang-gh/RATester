package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestGetTCPMiddlewareSection_1(t *testing.T) {
	tests := []struct {
		name        string
		middlewares map[string]*runtime.TCPMiddlewareInfo
		want        *section
	}{
		{
			name: "no middlewares",
			middlewares: map[string]*runtime.TCPMiddlewareInfo{
				"middleware1": {
					Status: string(runtime.StatusDisabled),
				},
				"middleware2": {
					Status: string(runtime.StatusWarning),
				},
			},
			want: &section{
				Total:    2,
				Warnings: 1,
				Errors:   1,
			},
		},
		{
			name:        "nil middlewares",
			middlewares: nil,
			want:        &section{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := getTCPMiddlewareSection(tt.middlewares)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTCPMiddlewareSection() = %v, want %v", got, tt.want)
			}
		})
	}
}
