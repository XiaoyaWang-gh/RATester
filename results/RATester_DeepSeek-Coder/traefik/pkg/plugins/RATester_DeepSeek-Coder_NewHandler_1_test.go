package plugins

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewHandler_1(t *testing.T) {
	type args struct {
		ctx  context.Context
		next http.Handler
	}
	tests := []struct {
		name    string
		m       *YaegiMiddleware
		args    args
		want    http.Handler
		wantErr bool
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

			got, err := tt.m.NewHandler(tt.args.ctx, tt.args.next)
			if (err != nil) != tt.wantErr {
				t.Errorf("YaegiMiddleware.NewHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YaegiMiddleware.NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
