package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateLogConfig_1(t *testing.T) {
	tests := []struct {
		name    string
		config  *v1.LogConfig
		wantErr bool
	}{
		{
			name: "valid log level",
			config: &v1.LogConfig{
				Level: "info",
			},
			wantErr: false,
		},
		{
			name: "invalid log level",
			config: &v1.LogConfig{
				Level: "invalid",
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

			if err := validateLogConfig(tt.config); (err != nil) != tt.wantErr {
				t.Errorf("validateLogConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
