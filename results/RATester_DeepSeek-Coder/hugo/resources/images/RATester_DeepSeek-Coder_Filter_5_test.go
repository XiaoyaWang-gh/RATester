package images

import (
	"fmt"
	"image"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestFilter_5(t *testing.T) {
	type args struct {
		src     image.Image
		filters []gift.Filter
	}
	tests := []struct {
		name    string
		p       *ImageProcessor
		args    args
		want    image.Image
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

			got, err := tt.p.Filter(tt.args.src, tt.args.filters...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImageProcessor.Filter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImageProcessor.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
