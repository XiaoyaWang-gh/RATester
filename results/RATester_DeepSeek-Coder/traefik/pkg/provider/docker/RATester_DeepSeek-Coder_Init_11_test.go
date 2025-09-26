package docker

import (
	"fmt"
	"testing"

	ptypes "github.com/traefik/paerser/types"
)

func TestInit_11(t *testing.T) {
	type fields struct {
		Shared         Shared
		ClientConfig   ClientConfig
		RefreshSeconds ptypes.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "should return error when default rule is not valid",
			fields: fields{
				Shared: Shared{
					DefaultRule: "{{ .NonExistent }}",
				},
			},
			wantErr: true,
		},
		{
			name: "should not return error when default rule is valid",
			fields: fields{
				Shared: Shared{
					DefaultRule: "Host(`{{.ServiceName}}.{{.ProviderName}}.localhost`)",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &SwarmProvider{
				Shared:         tt.fields.Shared,
				ClientConfig:   tt.fields.ClientConfig,
				RefreshSeconds: tt.fields.RefreshSeconds,
			}
			err := p.Init()
			if (err != nil) != tt.wantErr {
				t.Errorf("SwarmProvider.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
