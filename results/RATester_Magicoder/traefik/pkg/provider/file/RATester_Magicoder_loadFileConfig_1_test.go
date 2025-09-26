package file

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestloadFileConfig_1(t *testing.T) {
	ctx := context.Background()
	provider := &Provider{
		Directory: "/path/to/directory",
	}

	tests := []struct {
		name          string
		ctx           context.Context
		filename      string
		parseTemplate bool
		want          *dynamic.Configuration
		wantErr       bool
	}{
		{
			name:          "test case 1",
			ctx:           ctx,
			filename:      "/path/to/file",
			parseTemplate: true,
			want: &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"router1": {
							Rule: "Host(`example.com`)",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name:          "test case 2",
			ctx:           ctx,
			filename:      "/path/to/file",
			parseTemplate: false,
			want: &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"router2": {
							Rule: "Host(`example.com`)",
						},
					},
				},
			},
			wantErr: false,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := provider.loadFileConfig(tt.ctx, tt.filename, tt.parseTemplate)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadFileConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadFileConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
