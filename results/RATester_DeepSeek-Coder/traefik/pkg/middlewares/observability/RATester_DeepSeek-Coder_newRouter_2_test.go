package observability

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewRouter_2(t *testing.T) {
	type args struct {
		ctx        context.Context
		router     string
		routerRule string
		service    string
		next       http.Handler
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

			if got := newRouter(tt.args.ctx, tt.args.router, tt.args.routerRule, tt.args.service, tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
