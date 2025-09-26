package captcha

import (
	"fmt"
	"testing"
)

func TestDrawCircle_1(t *testing.T) {
	type args struct {
		x        int
		y        int
		radius   int
		colorIdx uint8
	}
	tests := []struct {
		name string
		m    *Image
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

			tt.m.drawCircle(tt.args.x, tt.args.y, tt.args.radius, tt.args.colorIdx)
		})
	}
}
