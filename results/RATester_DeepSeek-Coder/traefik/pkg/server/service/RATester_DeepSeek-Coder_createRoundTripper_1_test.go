package service

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestCreateRoundTripper_1(t *testing.T) {
	type args struct {
		cfg *dynamic.ServersTransport
	}
	tests := []struct {
		name    string
		args    args
		want    http.RoundTripper
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

			r := &RoundTripperManager{}
			got, err := r.createRoundTripper(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("RoundTripperManager.createRoundTripper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoundTripperManager.createRoundTripper() = %v, want %v", got, tt.want)
			}
		})
	}
}
