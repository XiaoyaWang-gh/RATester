package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/plugins"
)

func TestInitPlugins_1(t *testing.T) {
	type args struct {
		staticCfg *static.Configuration
	}
	tests := []struct {
		name    string
		args    args
		want    *plugins.Client
		want1   map[string]plugins.Descriptor
		want2   map[string]plugins.LocalDescriptor
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

			got, got1, got2, err := initPlugins(tt.args.staticCfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("initPlugins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initPlugins() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("initPlugins() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("initPlugins() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
