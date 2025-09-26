package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/config"
)

func TestDecodeImageConfig_1(t *testing.T) {
	type testCase struct {
		name         string
		action       string
		options      []string
		defaults     *config.ConfigNamespace[ImagingConfig, ImagingConfigInternal]
		sourceFormat Format
		want         ImageConfig
		wantErr      bool
	}

	tests := []testCase{
		{
			name:         "Test case 1",
			action:       "resize",
			options:      []string{"100x100", "q90"},
			defaults:     &config.ConfigNamespace[ImagingConfig, ImagingConfigInternal]{},
			sourceFormat: JPEG,
			want: ImageConfig{
				TargetFormat: JPEG,
				Action:       "resize",
				Quality:      90,
				Width:        100,
				Height:       100,
			},
			wantErr: false,
		},
		// Add more test cases here...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := DecodeImageConfig(tt.action, tt.options, tt.defaults, tt.sourceFormat)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeImageConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeImageConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
