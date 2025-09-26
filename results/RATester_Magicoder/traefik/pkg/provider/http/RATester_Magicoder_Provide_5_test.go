package http

import (
	"fmt"
	"net/http"
	"testing"

	ptypes "github.com/traefik/paerser/types"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/safe"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestProvide_5(t *testing.T) {
	type fields struct {
		Endpoint     string
		PollInterval ptypes.Duration
		PollTimeout  ptypes.Duration
		Headers      map[string]string
		TLS          *types.ClientTLS
		httpClient   *http.Client
	}
	type args struct {
		configurationChan chan<- dynamic.Message
		pool              *safe.Pool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
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

			p := &Provider{
				Endpoint:     tt.fields.Endpoint,
				PollInterval: tt.fields.PollInterval,
				PollTimeout:  tt.fields.PollTimeout,
				Headers:      tt.fields.Headers,
				TLS:          tt.fields.TLS,
				httpClient:   tt.fields.httpClient,
			}
			if err := p.Provide(tt.args.configurationChan, tt.args.pool); (err != nil) != tt.wantErr {
				t.Errorf("Provide() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
