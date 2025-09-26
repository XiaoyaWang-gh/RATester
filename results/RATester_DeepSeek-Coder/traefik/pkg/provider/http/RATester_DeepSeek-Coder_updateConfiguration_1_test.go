package http

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestUpdateConfiguration_1(t *testing.T) {
	type args struct {
		configurationChan chan<- dynamic.Message
	}
	tests := []struct {
		name    string
		p       *Provider
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

			if err := tt.p.updateConfiguration(tt.args.configurationChan); (err != nil) != tt.wantErr {
				t.Errorf("Provider.updateConfiguration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
