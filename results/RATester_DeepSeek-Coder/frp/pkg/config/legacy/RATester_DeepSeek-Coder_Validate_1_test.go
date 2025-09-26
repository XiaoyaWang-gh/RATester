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
				ServerAddr: "127.0.0.1",
				ServerPort: 7000,
				Protocol:   "tcp",
			},
			wantErr: false,
		},
		{
			name: "invalid server address",
			cfg: &ClientCommonConf{
				ServerAddr: "invalid",
				ServerPort: 7000,
				Protocol:   "tcp",
			},
			wantErr: true,
		},
		{
			name: "invalid server port",
			cfg: &ClientCommonConf{
				ServerAddr: "127.0.0.1",
				ServerPort: -1,
				Protocol:   "tcp",
			},
			wantErr: true,
		},
		{
			name: "invalid protocol",
			cfg: &ClientCommonConf{
				ServerAddr: "127.0.0.1",
				ServerPort: 7000,
				Protocol:   "invalid",
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

			err := tt.cfg.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
