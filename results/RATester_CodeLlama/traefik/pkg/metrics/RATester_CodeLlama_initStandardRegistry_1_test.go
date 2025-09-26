package metrics

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestInitStandardRegistry_1(t *testing.T) {
	type args struct {
		config *types.Prometheus
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				config: &types.Prometheus{
					Buckets:              []float64{0.1, 0.3, 1.2, 5.0},
					AddEntryPointsLabels: true,
					AddRoutersLabels:     true,
					AddServicesLabels:    true,
					EntryPoint:           "test",
					ManualRouting:        true,
					HeaderLabels:         map[string]string{"test": "test"},
				},
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

			initStandardRegistry(tt.args.config)
		})
	}
}
