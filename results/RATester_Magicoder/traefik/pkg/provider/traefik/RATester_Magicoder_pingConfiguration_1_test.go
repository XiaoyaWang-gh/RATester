package traefik

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestPingConfiguration_1(t *testing.T) {
	type args struct {
		cfg *dynamic.Configuration
	}
	tests := []struct {
		name string
		args args
		want *dynamic.Configuration
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

			i := &Provider{}
			i.pingConfiguration(tt.args.cfg)
		})
	}
}
