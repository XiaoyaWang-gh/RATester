package traefik

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestInit_19(t *testing.T) {
	type fields struct {
		staticCfg static.Configuration
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "should return no error when Init is called",
			fields: fields{
				staticCfg: static.Configuration{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			i := &Provider{
				staticCfg: tt.fields.staticCfg,
			}
			if err := i.Init(); (err != nil) != tt.wantErr {
				t.Errorf("Provider.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
