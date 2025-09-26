package images

import (
	"fmt"
	"image/color"
	"reflect"
	"testing"
)

func TestColorGoToColor_1(t *testing.T) {
	type args struct {
		c color.Color
	}
	tests := []struct {
		name string
		args args
		want Color
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

			if got := ColorGoToColor(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ColorGoToColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
