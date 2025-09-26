package captcha

import (
	"fmt"
	"testing"
)

func TestDrawDigit_1(t *testing.T) {
	type args struct {
		digit []byte
		x     int
		y     int
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

			tt.m.drawDigit(tt.args.digit, tt.args.x, tt.args.y)
		})
	}
}
