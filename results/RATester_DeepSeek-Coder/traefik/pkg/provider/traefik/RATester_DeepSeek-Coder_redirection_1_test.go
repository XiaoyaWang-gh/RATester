package traefik

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestRedirection_1(t *testing.T) {
	type args struct {
		ctx context.Context
		cfg *dynamic.Configuration
	}

	tests := []struct {
		name string
		args args
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
			i.redirection(tt.args.ctx, tt.args.cfg)
		})
	}
}
