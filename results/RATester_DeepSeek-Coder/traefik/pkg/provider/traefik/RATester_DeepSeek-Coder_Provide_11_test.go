package traefik

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/safe"
)

func TestProvide_11(t *testing.T) {
	type args struct {
		configurationChan chan<- dynamic.Message
		pool              *safe.Pool
	}

	tests := []struct {
		name    string
		i       *Provider
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

			err := tt.i.Provide(tt.args.configurationChan, tt.args.pool)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provide() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
