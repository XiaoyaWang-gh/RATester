package captcha

import (
	"fmt"
	"testing"
)

func TestDrawHorizLine_1(t *testing.T) {
	type args struct {
		fromX    int
		toX      int
		y        int
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

			tt.m.drawHorizLine(tt.args.fromX, tt.args.toX, tt.args.y, tt.args.colorIdx)
		})
	}
}
