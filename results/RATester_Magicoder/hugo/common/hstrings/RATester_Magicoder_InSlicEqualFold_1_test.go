package hstrings

import (
	"fmt"
	"testing"
)

func TestInSlicEqualFold_1(t *testing.T) {
	type args struct {
		arr []string
		el  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				arr: []string{"Hello", "World", "Golang"},
				el:  "hello",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				arr: []string{"Hello", "World", "Golang"},
				el:  "python",
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				arr: []string{"Hello", "World", "Golang"},
				el:  "Golang",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := InSlicEqualFold(tt.args.arr, tt.args.el); got != tt.want {
				t.Errorf("InSlicEqualFold() = %v, want %v", got, tt.want)
			}
		})
	}
}
