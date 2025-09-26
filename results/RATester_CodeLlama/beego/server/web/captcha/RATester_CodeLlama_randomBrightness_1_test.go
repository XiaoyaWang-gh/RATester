package captcha

import (
	"fmt"
	"image/color"
	"reflect"
	"testing"
)

func TestRandomBrightness_1(t *testing.T) {
	type args struct {
		c   color.RGBA
		max uint8
	}
	tests := []struct {
		name string
		args args
		want color.RGBA
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

			if got := randomBrightness(tt.args.c, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("randomBrightness() = %v, want %v", got, tt.want)
			}
		})
	}
}
