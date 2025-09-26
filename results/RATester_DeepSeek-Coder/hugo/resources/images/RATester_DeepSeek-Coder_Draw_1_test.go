package images

import (
	"fmt"
	"image"
	"image/draw"
	"testing"

	"github.com/disintegration/gift"
)

func TestDraw_1(t *testing.T) {
	type args struct {
		dst     draw.Image
		src     image.Image
		options *gift.Options
	}
	tests := []struct {
		name string
		f    textFilter
		args args
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

			f := textFilter{
				text:        tt.f.text,
				color:       tt.f.color,
				x:           tt.f.x,
				y:           tt.f.y,
				size:        tt.f.size,
				linespacing: tt.f.linespacing,
				fontSource:  tt.f.fontSource,
			}
			f.Draw(tt.args.dst, tt.args.src, tt.args.options)
		})
	}
}
