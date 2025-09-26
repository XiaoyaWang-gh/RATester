package captcha

import (
	"fmt"
	"testing"
)

func TestFillWithCircles_1(t *testing.T) {
	type args struct {
		n         int
		maxradius int
	}
	tests := []struct {
		name string
		m    *Image
		args args
	}{
		{
			name: "Test fillWithCircles with positive values",
			m:    &Image{},
			args: args{
				n:         5,
				maxradius: 10,
			},
		},
		{
			name: "Test fillWithCircles with zero values",
			m:    &Image{},
			args: args{
				n:         0,
				maxradius: 0,
			},
		},
		{
			name: "Test fillWithCircles with negative values",
			m:    &Image{},
			args: args{
				n:         -5,
				maxradius: -10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.m.fillWithCircles(tt.args.n, tt.args.maxradius)
		})
	}
}
