package captcha

import (
	"fmt"
	"testing"
)

func TestDistort_1(t *testing.T) {
	type args struct {
		amplude float64
		period  float64
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

			tt.m.distort(tt.args.amplude, tt.args.period)
		})
	}
}
