package auth

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestNewForward_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		next   http.Handler
		config dynamic.ForwardAuth
		name   string
	}
	tests := []struct {
		name    string
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

			got, err := NewForward(tt.args.ctx, tt.args.next, tt.args.config, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewForward() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewForward() = %v, want %v", got, tt.want)
			}
		})
	}
}
