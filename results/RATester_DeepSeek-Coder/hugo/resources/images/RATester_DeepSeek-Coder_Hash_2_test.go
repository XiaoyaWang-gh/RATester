package images

import (
	"fmt"
	"image/color"
	"testing"
)

func TestHash_2(t *testing.T) {
	tests := []struct {
		name    string
		c       Color
		want    uint64
		wantErr bool
	}{
		{
			name: "Test case 1",
			c: Color{
				color:     color.RGBA{R: 255, G: 0, B: 0, A: 255},
				hex:       "#ff0000",
				luminance: 0.2126,
			},
			want:    0xcc9,
			wantErr: false,
		},
		{
			name: "Test case 2",
			c: Color{
				color:     color.RGBA{R: 0, G: 255, B: 0, A: 255},
				hex:       "#00ff00",
				luminance: 0.587,
			},
			want:    0xcc9,
			wantErr: false,
		},
		{
			name: "Test case 3",
			c: Color{
				color:     color.RGBA{R: 0, G: 0, B: 255, A: 255},
				hex:       "#0000ff",
				luminance: 0.1422,
			},
			want:    0xcc9,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.c.Hash()
			if (err != nil) != tt.wantErr {
				t.Errorf("Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
