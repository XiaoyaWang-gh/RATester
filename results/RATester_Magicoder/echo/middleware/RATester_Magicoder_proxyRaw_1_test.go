package middleware

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestProxyRaw_1(t *testing.T) {
	type args struct {
		t *ProxyTarget
		c echo.Context
	}
	tests := []struct {
		name string
		args args
		want http.Handler
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

			if got := proxyRaw(tt.args.t, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("proxyRaw() = %v, want %v", got, tt.want)
			}
		})
	}
}
