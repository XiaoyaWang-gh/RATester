package compare

import (
	"fmt"
	"testing"
)

func TestCompareFold_1(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				s: "Hello",
				t: "hello",
			},
			want: 0,
		},
		{
			name: "Test Case 2",
			args: args{
				s: "HELLO",
				t: "hello",
			},
			want: -1,
		},
		{
			name: "Test Case 3",
			args: args{
				s: "hello",
				t: "HELLO",
			},
			want: 1,
		},
		{
			name: "Test Case 4",
			args: args{
				s: "HELLO",
				t: "HELLO",
			},
			want: 0,
		},
		{
			name: "Test Case 5",
			args: args{
				s: "HEllO",
				t: "HELLO",
			},
			want: -1,
		},
		{
			name: "Test Case 6",
			args: args{
				s: "HELLO",
				t: "HEllO",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := compareFold(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("compareFold() = %v, want %v", got, tt.want)
			}
		})
	}
}
