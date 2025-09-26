package images

import (
	"fmt"
	"image"
	"reflect"
	"testing"
)

func TestResize_2(t *testing.T) {
	type args struct {
		img    image.Image
		width  uint
		height uint
	}
	tests := []struct {
		name string
		args args
		want image.Image
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

			r := imagingResizer{}
			if got := r.Resize(tt.args.img, tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resize() = %v, want %v", got, tt.want)
			}
		})
	}
}
