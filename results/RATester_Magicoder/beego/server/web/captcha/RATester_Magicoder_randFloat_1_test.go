package captcha

import (
	"fmt"
	"testing"
)

func TestrandFloat_1(t *testing.T) {
	type args struct {
		from float64
		to   float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test case 1",
			args: args{
				from: 1.0,
				to:   10.0,
			},
			want: 5.0, // This is a placeholder, replace with the expected result
		},
		{
			name: "Test case 2",
			args: args{
				from: 0.0,
				to:   1.0,
			},
			want: 0.5, // This is a placeholder, replace with the expected result
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

			if got := randFloat(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("randFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
