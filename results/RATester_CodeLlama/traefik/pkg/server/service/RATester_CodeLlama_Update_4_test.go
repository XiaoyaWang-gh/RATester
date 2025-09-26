package service

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestUpdate_4(t *testing.T) {
	type args struct {
		newConfigs map[string]*dynamic.ServersTransport
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				newConfigs: map[string]*dynamic.ServersTransport{
					"test": {
						ServerName: "test",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := &RoundTripperManager{
				roundTrippers: map[string]http.RoundTripper{},
				configs:       map[string]*dynamic.ServersTransport{},
			}
			r.Update(tt.args.newConfigs)
		})
	}
}
