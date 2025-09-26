package images

import (
	"fmt"
	"image"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestFiltersFromConfig_1(t *testing.T) {
	type args struct {
		src  image.Image
		conf ImageConfig
	}
	tests := []struct {
		name    string
		args    args
		want    []gift.Filter
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

			p := &ImageProcessor{}
			got, err := p.FiltersFromConfig(tt.args.src, tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImageProcessor.FiltersFromConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImageProcessor.FiltersFromConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
