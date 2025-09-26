package auth

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestNewDigest_1(t *testing.T) {
	type args struct {
		ctx        context.Context
		next       http.Handler
		authConfig dynamic.DigestAuth
		name       string
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

			got, err := NewDigest(tt.args.ctx, tt.args.next, tt.args.authConfig, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDigest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDigest() = %v, want %v", got, tt.want)
			}
		})
	}
}
