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
		// TODO: Add test cases.
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
