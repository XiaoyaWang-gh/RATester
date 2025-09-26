package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestgetTCPMiddlewareSection_1(t *testing.T) {
	tests := []struct {
		name        string
		middlewares map[string]*runtime.TCPMiddlewareInfo
		want        *section
	}{
		{
			name:        "no middlewares",
			middlewares: map[string]*runtime.TCPMiddlewareInfo{},
			want: &section{
				Total:    0,
				Warnings: 0,
				Errors:   0,
			},
		},
		{
			name: "one warning and one error",
			middlewares: map[string]*runtime.TCPMiddlewareInfo{
				"warning": {
					Status: string(runtime.StatusWarning),
				},
				"error": {
					Status: string(runtime.StatusDisabled),
				},
			},
			want: &section{
				Total:    2,
				Warnings: 1,
				Errors:   1,
			},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getTCPMiddlewareSection(tt.middlewares); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTCPMiddlewareSection() = %v, want %v", got, tt.want)
			}
		})
	}
}
