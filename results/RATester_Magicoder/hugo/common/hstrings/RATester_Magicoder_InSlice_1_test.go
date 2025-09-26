package hstrings

import (
	"fmt"
	"testing"
)

func TestInSlice_1(t *testing.T) {
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
			name: "Test case 1",
			args: args{
				arr: []string{"a", "b", "c"},
				el:  "b",
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				arr: []string{"a", "b", "c"},
				el:  "d",
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				arr: []string{},
				el:  "a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := InSlice(tt.args.arr, tt.args.el); got != tt.want {
				t.Errorf("InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
