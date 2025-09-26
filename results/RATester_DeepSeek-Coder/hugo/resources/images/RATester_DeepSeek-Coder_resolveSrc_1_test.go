package images

import (
	"fmt"
	"image"
	"reflect"
	"testing"
)

func TestResolveSrc_1(t *testing.T) {
	type args struct {
		src          image.Image
		targetFormat Format
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

			p := &ImageProcessor{}
			if got := p.resolveSrc(tt.args.src, tt.args.targetFormat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImageProcessor.resolveSrc() = %v, want %v", got, tt.want)
			}
		})
	}
}
