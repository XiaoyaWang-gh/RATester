package images

import (
	"fmt"
	"math"
	"testing"
)

func Testsinc_1(t *testing.T) {
	type args struct {
		x float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "Test case 1",
			args: args{
				x: 0,
			},
			want: 1,
		},
		{
			name: "Test case 2",
			args: args{
				x: 1,
			},
			want: float32(math.Sin(math.Pi*1) / (math.Pi * 1)),
		},
		{
			name: "Test case 3",
			args: args{
				x: -1,
			},
			want: float32(math.Sin(math.Pi*-1) / (math.Pi * -1)),
		},
		{
			name: "Test case 4",
			args: args{
				x: 10,
			},
			want: float32(math.Sin(math.Pi*10) / (math.Pi * 10)),
		},
		{
			name: "Test case 5",
			args: args{
				x: -10,
			},
			want: float32(math.Sin(math.Pi*-10) / (math.Pi * -10)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := sinc(tt.args.x); got != tt.want {
				t.Errorf("sinc() = %v, want %v", got, tt.want)
			}
		})
	}
}
