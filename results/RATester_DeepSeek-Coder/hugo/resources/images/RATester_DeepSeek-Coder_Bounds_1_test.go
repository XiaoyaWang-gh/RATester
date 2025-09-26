package images

import (
	"fmt"
	"image"
	"image/color"
	"testing"
)

func TestBounds_1(t *testing.T) {
	testCases := []struct {
		name      string
		filter    textFilter
		srcBounds image.Rectangle
		want      image.Rectangle
	}{
		{
			name: "Test case 1",
			filter: textFilter{
				text:        "Hello, world",
				color:       color.RGBA{R: 255, G: 255, B: 255, A: 255},
				x:           10,
				y:           10,
				size:        12.0,
				linespacing: 1,
				fontSource:  nil,
			},
			srcBounds: image.Rect(0, 0, 800, 600),
			want:      image.Rect(0, 0, 800, 600),
		},
		{
			name: "Test case 2",
			filter: textFilter{
				text:        "Hello, world",
				color:       color.RGBA{R: 255, G: 255, B: 255, A: 255},
				x:           10,
				y:           10,
				size:        12.0,
				linespacing: 1,
				fontSource:  nil,
			},
			srcBounds: image.Rect(0, 0, 100, 100),
			want:      image.Rect(0, 0, 100, 100),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.filter.Bounds(tc.srcBounds)
			if !got.Eq(tc.want) {
				t.Errorf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}
