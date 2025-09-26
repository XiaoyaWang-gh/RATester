package observability

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/containous/alice"
)

func TestWrapRouterHandler_2(t *testing.T) {
	type args struct {
		ctx        context.Context
		router     string
		routerRule string
		service    string
	}
	tests := []struct {
		name string
		args args
		want alice.Constructor
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

			if got := WrapRouterHandler(tt.args.ctx, tt.args.router, tt.args.routerRule, tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapRouterHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
