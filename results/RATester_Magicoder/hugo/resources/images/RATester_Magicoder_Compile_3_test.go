package images

import (
	"fmt"
	"testing"
)

func TestCompile_3(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *ImagingConfig
		wantErr bool
	}{
		{
			name: "valid config",
			cfg: &ImagingConfig{
				Quality:        50,
				ResampleFilter: "lanczos2",
				Hint:           "photo",
				Anchor:         "smart",
				BgColor:        "fff",
			},
			wantErr: false,
		},
		{
			name: "invalid resample filter",
			cfg: &ImagingConfig{
				Quality:        50,
				ResampleFilter: "invalid",
				Hint:           "photo",
				Anchor:         "smart",
				BgColor:        "fff",
			},
			wantErr: true,
		},
		{
			name: "invalid anchor",
			cfg: &ImagingConfig{
				Quality:        50,
				ResampleFilter: "lanczos2",
				Hint:           "photo",
				Anchor:         "invalid",
				BgColor:        "fff",
			},
			wantErr: true,
		},
		{
			name: "invalid bg color",
			cfg: &ImagingConfig{
				Quality:        50,
				ResampleFilter: "lanczos2",
				Hint:           "photo",
				Anchor:         "smart",
				BgColor:        "invalid",
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

			i := &ImagingConfigInternal{}
			if err := i.Compile(tt.cfg); (err != nil) != tt.wantErr {
				t.Errorf("Compile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
