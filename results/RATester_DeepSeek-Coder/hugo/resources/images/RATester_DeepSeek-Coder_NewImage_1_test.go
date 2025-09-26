package images

import (
	"fmt"
	"image"
	"reflect"
	"testing"
)

func TestNewImage_1(t *testing.T) {
	type args struct {
		f    Format
		proc *ImageProcessor
		img  image.Image
		s    Spec
	}
	tests := []struct {
		name string
		args args
		want *Image
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

			if got := NewImage(tt.args.f, tt.args.proc, tt.args.img, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
