package docker

import (
	"context"
	"fmt"
	"testing"
)

func TestKeepContainer_1(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name      string
		container dockerData
		want      bool
	}{
		{
			name: "Container with disabled label",
			container: dockerData{
				ExtraConf: configuration{
					Enable: false,
				},
			},
			want: false,
		},
		{
			name: "Container with unhealthy health status",
			container: dockerData{
				Health: "unhealthy",
				ExtraConf: configuration{
					Enable: true,
				},
			},
			want: false,
		},
		{
			name: "Container with healthy health status",
			container: dockerData{
				Health: "healthy",
				ExtraConf: configuration{
					Enable: true,
				},
			},
			want: true,
		},
		{
			name: "Container with invalid constraints",
			container: dockerData{
				Labels: map[string]string{
					"constraints": "invalid",
				},
				ExtraConf: configuration{
					Enable: true,
				},
			},
			want: false,
		},
		{
			name: "Container with valid constraints",
			container: dockerData{
				Labels: map[string]string{
					"constraints": "label==value",
				},
				ExtraConf: configuration{
					Enable: true,
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			dynConfBuilder := &DynConfBuilder{}
			if got := dynConfBuilder.keepContainer(ctx, tt.container); got != tt.want {
				t.Errorf("keepContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}
