package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestGetHTTPMiddlewareSection_1(t *testing.T) {
	type testCase struct {
		name        string
		middlewares map[string]*runtime.MiddlewareInfo
		want        *section
	}

	testCases := []testCase{
		{
			name:        "no middlewares",
			middlewares: map[string]*runtime.MiddlewareInfo{},
			want: &section{
				Total:    0,
				Warnings: 0,
				Errors:   0,
			},
		},
		{
			name: "one warning and one error",
			middlewares: map[string]*runtime.MiddlewareInfo{
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

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := getHTTPMiddlewareSection(tc.middlewares)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
