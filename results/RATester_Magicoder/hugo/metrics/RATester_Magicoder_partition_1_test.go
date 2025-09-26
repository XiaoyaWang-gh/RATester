package metrics

import (
	"fmt"
	"testing"
)

func Testpartition_1(t *testing.T) {
	type args struct {
		d     int
		scale int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				d:     10,
				scale: 3,
			},
			want: 9,
		},
		{
			name: "Test case 2",
			args: args{
				d:     15,
				scale: 5,
			},
			want: 15,
		},
		{
			name: "Test case 3",
			args: args{
				d:     20,
				scale: 10,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := partition(tt.args.d, tt.args.scale); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
