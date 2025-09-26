package images

import (
	"fmt"
	"image/color"
	"reflect"
	"testing"
)

func TestAddColorToPalette_1(t *testing.T) {
	type args struct {
		c color.Color
		p color.Palette
	}
	tests := []struct {
		name string
		args args
		want color.Palette
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

			if got := AddColorToPalette(tt.args.c, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddColorToPalette() = %v, want %v", got, tt.want)
			}
		})
	}
}
