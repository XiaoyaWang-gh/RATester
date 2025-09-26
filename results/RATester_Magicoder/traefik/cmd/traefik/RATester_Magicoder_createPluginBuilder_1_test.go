package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/plugins"
)

func TestCreatePluginBuilder_1(t *testing.T) {
	type args struct {
		staticConfiguration *static.Configuration
	}
	tests := []struct {
		name    string
		args    args
		want    *plugins.Builder
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				staticConfiguration: &static.Configuration{
					// Add your configuration here
				},
			},
			want: &plugins.Builder{
				// Add your expected builder here
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				staticConfiguration: &static.Configuration{
					// Add your configuration here
				},
			},
			want: &plugins.Builder{
				// Add your expected builder here
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

			got, err := createPluginBuilder(tt.args.staticConfiguration)
			if (err != nil) != tt.wantErr {
				t.Errorf("createPluginBuilder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createPluginBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
