package images

import (
	"fmt"
	"testing"
)

func Testabsf32_1(t *testing.T) {
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
				x: 1.0,
			},
			want: 1.0,
		},
		{
			name: "Test case 2",
			args: args{
				x: -1.0,
			},
			want: 1.0,
		},
		{
			name: "Test case 3",
			args: args{
				x: 0.0,
			},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := absf32(tt.args.x); got != tt.want {
				t.Errorf("absf32() = %v, want %v", got, tt.want)
			}
		})
	}
}
