package captcha

import (
	"fmt"
	"testing"
)

func Testmin3_1(t *testing.T) {
	type args struct {
		x uint8
		y uint8
		z uint8
	}
	tests := []struct {
		name  string
		args  args
		wantM uint8
	}{
		{
			name: "Test Case 1",
			args: args{
				x: 1,
				y: 2,
				z: 3,
			},
			wantM: 1,
		},
		{
			name: "Test Case 2",
			args: args{
				x: 4,
				y: 5,
				z: 6,
			},
			wantM: 4,
		},
		{
			name: "Test Case 3",
			args: args{
				x: 7,
				y: 8,
				z: 9,
			},
			wantM: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if gotM := min3(tt.args.x, tt.args.y, tt.args.z); gotM != tt.wantM {
				t.Errorf("min3() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}
