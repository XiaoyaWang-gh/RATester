package images

import (
	"fmt"
	"image"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestSmartCrop_1(t *testing.T) {
	type args struct {
		img    image.Image
		width  int
		height int
		filter gift.Resampling
	}
	tests := []struct {
		name    string
		p       *ImageProcessor
		args    args
		want    image.Rectangle
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

			got, err := tt.p.smartCrop(tt.args.img, tt.args.width, tt.args.height, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImageProcessor.smartCrop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImageProcessor.smartCrop() = %v, want %v", got, tt.want)
			}
		})
	}
}
