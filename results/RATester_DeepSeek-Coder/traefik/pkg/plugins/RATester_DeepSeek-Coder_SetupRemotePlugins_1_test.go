package plugins

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestSetupRemotePlugins_1(t *testing.T) {
	type args struct {
		client  *Client
		plugins map[string]Descriptor
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with valid client and plugins",
			args: args{
				client: &Client{
					HTTPClient: &http.Client{},
					baseURL:    &url.URL{},
				},
				plugins: map[string]Descriptor{
					"plugin1": {
						ModuleName: "validModuleName",
						Version:    "validVersion",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Test with invalid client",
			args: args{
				client: &Client{
					HTTPClient: nil,
					baseURL:    &url.URL{},
				},
				plugins: map[string]Descriptor{
					"plugin1": {
						ModuleName: "validModuleName",
						Version:    "validVersion",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Test with invalid plugins",
			args: args{
				client: &Client{
					HTTPClient: &http.Client{},
					baseURL:    &url.URL{},
				},
				plugins: map[string]Descriptor{
					"plugin1": {
						ModuleName: "",
						Version:    "validVersion",
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := SetupRemotePlugins(tt.args.client, tt.args.plugins); (err != nil) != tt.wantErr {
				t.Errorf("SetupRemotePlugins() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
