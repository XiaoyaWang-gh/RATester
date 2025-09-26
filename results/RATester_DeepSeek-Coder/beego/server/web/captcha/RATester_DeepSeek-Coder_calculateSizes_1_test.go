package captcha

import (
	"fmt"
	"testing"
)

func TestCalculateSizes_1(t *testing.T) {
	type args struct {
		width  int
		height int
		ncount int
	}
	tests := []struct {
		name string
		args args
		want Image
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

			m := &Image{}
			m.calculateSizes(tt.args.width, tt.args.height, tt.args.ncount)
			if m.numWidth != tt.want.numWidth || m.numHeight != tt.want.numHeight || m.dotSize != tt.want.dotSize {
				t.Errorf("calculateSizes() = %v, want %v", m, tt.want)
			}
		})
	}
}
