package images

import (
	"fmt"
	"testing"
)

func Testinit_30(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *ImagingConfig
		wantErr bool
	}{
		{
			name: "Quality less than 0",
			cfg: &ImagingConfig{
				Quality: -1,
			},
			wantErr: true,
		},
		{
			name: "Quality more than 100",
			cfg: &ImagingConfig{
				Quality: 101,
			},
			wantErr: true,
		},
		{
			name: "Valid Quality",
			cfg: &ImagingConfig{
				Quality: 50,
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

			if err := tt.cfg.init(); (err != nil) != tt.wantErr {
				t.Errorf("init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
