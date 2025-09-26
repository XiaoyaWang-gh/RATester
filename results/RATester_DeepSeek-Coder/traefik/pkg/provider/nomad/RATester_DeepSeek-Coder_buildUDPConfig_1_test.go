package nomad

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildUDPConfig_1(t *testing.T) {
	type args struct {
		i             item
		configuration *dynamic.UDPConfiguration
	}
	tests := []struct {
		name    string
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

			p := &Provider{}
			if err := p.buildUDPConfig(tt.args.i, tt.args.configuration); (err != nil) != tt.wantErr {
				t.Errorf("Provider.buildUDPConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
