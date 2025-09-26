package main

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/plugins"
)

func TestCheckUniquePluginNames_1(t *testing.T) {
	type args struct {
		e *static.Experimental
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "nil experimental",
			args: args{
				e: nil,
			},
			wantErr: false,
		},
		{
			name: "duplicate plugin names",
			args: args{
				e: &static.Experimental{
					Plugins: map[string]plugins.Descriptor{
						"plugin1": {},
					},
					LocalPlugins: map[string]plugins.LocalDescriptor{
						"plugin1": {},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "unique plugin names",
			args: args{
				e: &static.Experimental{
					Plugins: map[string]plugins.Descriptor{
						"plugin1": {},
					},
					LocalPlugins: map[string]plugins.LocalDescriptor{
						"plugin2": {},
					},
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

			if err := checkUniquePluginNames(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("checkUniquePluginNames() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
