package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bep/logg"
	"github.com/gohugoio/hugo/config"
)

func TestNewImageProcessor_1(t *testing.T) {
	type args struct {
		warnl logg.LevelLogger
		cfg   *config.ConfigNamespace[ImagingConfig, ImagingConfigInternal]
	}
	tests := []struct {
		name    string
		args    args
		want    *ImageProcessor
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := NewImageProcessor(tt.args.warnl, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewImageProcessor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImageProcessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
