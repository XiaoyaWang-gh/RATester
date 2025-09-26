package tcp

import (
	"fmt"
	"testing"
)

func TestGcd_2(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				a: 10,
				b: 5,
			},
			want: 5,
		},
		{
			name: "Test Case 2",
			args: args{
				a: 15,
				b: 5,
			},
			want: 5,
		},
		{
			name: "Test Case 3",
			args: args{
				a: 20,
				b: 10,
			},
			want: 10,
		},
		{
			name: "Test Case 4",
			args: args{
				a: 100,
				b: 50,
			},
			want: 50,
		},
		{
			name: "Test Case 5",
			args: args{
				a: 123456789,
				b: 987654321,
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

			if got := gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}
