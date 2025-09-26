package http

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	ptypes "github.com/traefik/paerser/types"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestFetchConfigurationData_1(t *testing.T) {
	type fields struct {
		Endpoint     string
		PollInterval ptypes.Duration
		PollTimeout  ptypes.Duration
		Headers      map[string]string
		TLS          *types.ClientTLS
		httpClient   *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
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
			got, err := p.fetchConfigurationData()
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.fetchConfigurationData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.fetchConfigurationData() = %v, want %v", got, tt.want)
			}
		})
	}
}
