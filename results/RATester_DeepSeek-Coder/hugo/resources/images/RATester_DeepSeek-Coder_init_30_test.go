package images

import (
	"fmt"
	"testing"
)

func TestInit_30(t *testing.T) {
	type testCase struct {
		name    string
		cfg     ImagingConfig
		wantErr bool
	}

	tests := []testCase{
		{
			name: "valid config",
			cfg: ImagingConfig{
				Quality: 50,
				BgColor: "fff",
			},
			wantErr: false,
		},
		{
			name: "invalid quality",
			cfg: ImagingConfig{
				Quality: 101,
				BgColor: "fff",
			},
			wantErr: true,
		},
		{
			name: "invalid color",
			cfg: ImagingConfig{
				Quality: 50,
				BgColor: "fffff",
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

			err := tt.cfg.init()
			if (err != nil) != tt.wantErr {
				t.Errorf("init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
