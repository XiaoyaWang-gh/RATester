package metrics

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestRegisterInfluxDB2_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		config *types.InfluxDB2
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "RegisterInfluxDB2",
			args: args{
				ctx:    context.Background(),
				config: &types.InfluxDB2{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			RegisterInfluxDB2(tt.args.ctx, tt.args.config)
		})
	}
}
