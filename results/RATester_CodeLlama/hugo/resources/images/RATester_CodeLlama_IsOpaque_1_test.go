package images

import (
	"fmt"
	"image"
	"testing"
)

func TestIsOpaque_1(t *testing.T) {
	type args struct {
		img image.Image
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1",
			args: args{
				img: image.NewRGBA(image.Rect(0, 0, 10, 10)),
			},
			want: true,
		},
		{
			name: "test case 2",
			args: args{
				img: image.NewRGBA(image.Rect(0, 0, 10, 10)),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsOpaque(tt.args.img); got != tt.want {
				t.Errorf("IsOpaque() = %v, want %v", got, tt.want)
			}
		})
	}
}
