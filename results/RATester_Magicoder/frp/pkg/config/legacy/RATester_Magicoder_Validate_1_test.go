package legacy

import (
	"fmt"
	"testing"
)

func TestValidate_1(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *ClientCommonConf
		wantErr bool
	}{
		{
			name: "valid config",
			cfg: &ClientCommonConf{
				HeartbeatTimeout:  100,
				HeartbeatInterval: 50,
			},
			wantErr: false,
		},
		{
			name: "invalid config",
			cfg: &ClientCommonConf{
				HeartbeatTimeout:  10,
				HeartbeatInterval: 50,
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

			if err := tt.cfg.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
