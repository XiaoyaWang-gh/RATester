package images

import (
	"fmt"
	"testing"
)

func TestInit_30(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *ImagingConfig
		wantErr bool
	}{
		{
			name: "valid",
			cfg: &ImagingConfig{
				Quality: 100,
			},
			wantErr: false,
		},
		{
			name: "invalid",
			cfg: &ImagingConfig{
				Quality: -1,
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

			if err := tt.cfg.init(); (err != nil) != tt.wantErr {
				t.Errorf("init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
