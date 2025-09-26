package metrics

import (
	"fmt"
	"testing"
)

func TesthowSimilarStrings_1(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				a: "Hello World",
				b: "Hello World",
			},
			want: 100,
		},
		{
			name: "Test Case 2",
			args: args{
				a: "Hello World",
				b: "Hello Universe",
			},
			want: 50,
		},
		{
			name: "Test Case 3",
			args: args{
				a: "Hello World",
				b: "Hello",
			},
			want: 100,
		},
		{
			name: "Test Case 4",
			args: args{
				a: "Hello World",
				b: "World Hello",
			},
			want: 100,
		},
		{
			name: "Test Case 5",
			args: args{
				a: "Hello World",
				b: "Hello",
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := howSimilarStrings(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("howSimilarStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
