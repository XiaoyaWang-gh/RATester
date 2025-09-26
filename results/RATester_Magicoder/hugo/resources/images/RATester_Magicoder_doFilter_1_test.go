package images

import (
	"fmt"
	"image"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestdoFilter_1(t *testing.T) {
	type args struct {
		src          image.Image
		targetFormat Format
		filters      []gift.Filter
	}
	tests := []struct {
		name    string
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

			p := &ImageProcessor{}
			got, err := p.doFilter(tt.args.src, tt.args.targetFormat, tt.args.filters...)
			if (err != nil) != tt.wantErr {
				t.Errorf("doFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
