package observability

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewService_1(t *testing.T) {
	type args struct {
		ctx     context.Context
		service string
		next    http.Handler
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

			if got := NewService(tt.args.ctx, tt.args.service, tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}
