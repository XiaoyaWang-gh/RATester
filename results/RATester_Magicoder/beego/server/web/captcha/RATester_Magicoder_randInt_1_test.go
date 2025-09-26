package captcha

import (
	"fmt"
	"testing"
)

func TestrandInt_1(t *testing.T) {
	type args struct {
		from int
		to   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				from: 1,
				to:   10,
			},
			want: 5,
		},
		{
			name: "Test case 2",
			args: args{
				from: 10,
				to:   100,
			},
			want: 50,
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

			if got := randInt(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("randInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
