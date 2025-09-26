package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateVisitorBaseConfig_1(t *testing.T) {
	tests := []struct {
		name    string
		config  *v1.VisitorBaseConfig
		wantErr bool
	}{
		{
			name: "valid config",
			config: &v1.VisitorBaseConfig{
				Name:       "validName",
				ServerName: "validServerName",
				BindPort:   8080,
			},
			wantErr: false,
		},
		{
			name: "missing name",
			config: &v1.VisitorBaseConfig{
				ServerName: "validServerName",
				BindPort:   8080,
			},
			wantErr: true,
		},
		{
			name: "missing server name",
			config: &v1.VisitorBaseConfig{
				Name:     "validName",
				BindPort: 8080,
			},
			wantErr: true,
		},
		{
			name: "missing bind port",
			config: &v1.VisitorBaseConfig{
				Name:       "validName",
				ServerName: "validServerName",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := validateVisitorBaseConfig(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateVisitorBaseConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
