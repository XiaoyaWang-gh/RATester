package consul

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestInit_20(t *testing.T) {
	tests := []struct {
		name      string
		namespace string
		token     string
		tls       *types.ClientTLS
		wantErr   bool
	}{
		{
			name:      "test",
			namespace: "test",
			token:     "test",
			tls:       &types.ClientTLS{},
			wantErr:   false,
		},
		{
			name:      "test",
			namespace: "*",
			token:     "test",
			tls:       &types.ClientTLS{},
			wantErr:   true,
		},
		{
			name:      "test",
			namespace: "test",
			token:     "test",
			tls:       nil,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Provider{
				name:      tt.name,
				namespace: tt.namespace,
				token:     tt.token,
				tls:       tt.tls,
			}
			err := p.Init()
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
