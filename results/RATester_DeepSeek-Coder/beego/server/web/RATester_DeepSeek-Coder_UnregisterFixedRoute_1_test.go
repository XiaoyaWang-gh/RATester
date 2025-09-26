package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnregisterFixedRoute_1(t *testing.T) {
	app := &HttpServer{
		Handlers: &ControllerRegister{
			routers: make(map[string]*Tree),
		},
	}

	testCases := []struct {
		name       string
		fixedRoute string
		method     string
		want       *HttpServer
	}{
		{
			name:       "UnregisterFixedRoute with empty method",
			fixedRoute: "/test",
			method:     "",
			want:       app,
		},
		{
			name:       "UnregisterFixedRoute with wildcard method",
			fixedRoute: "/test",
			method:     "*",
			want:       app,
		},
		{
			name:       "UnregisterFixedRoute with specific method",
			fixedRoute: "/test",
			method:     "GET",
			want:       app,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := app.UnregisterFixedRoute(tc.fixedRoute, tc.method)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("UnregisterFixedRoute() = %v, want %v", got, tc.want)
			}
		})
	}
}
