package compare

import (
	"fmt"
	"testing"
)

func TestProbablyEq_3(t *testing.T) {
	type args struct {
		v1 any
		v2 any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				v1: 1,
				v2: 1,
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				v1: 1,
				v2: 2,
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				v1: "test",
				v2: "test",
			},
			want: true,
		},
		{
			name: "Test case 4",
			args: args{
				v1: "test",
				v2: "not test",
			},
			want: false,
		},
		{
			name: "Test case 5",
			args: args{
				v1: []int{1, 2, 3},
				v2: []int{1, 2, 3},
			},
			want: true,
		},
		{
			name: "Test case 6",
			args: args{
				v1: []int{1, 2, 3},
				v2: []int{1, 2, 4},
			},
			want: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ProbablyEq(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("ProbablyEq() = %v, want %v", got, tt.want)
			}
		})
	}
}
