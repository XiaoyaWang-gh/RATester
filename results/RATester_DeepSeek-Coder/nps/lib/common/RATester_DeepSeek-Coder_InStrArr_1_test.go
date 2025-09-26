package common

import (
	"fmt"
	"testing"
)

func TestInStrArr_1(t *testing.T) {
	type args struct {
		arr []string
		val string
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
				val: "b",
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				arr: []string{"a", "b", "c"},
				val: "d",
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				arr: []string{},
				val: "d",
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

			if got := InStrArr(tt.args.arr, tt.args.val); got != tt.want {
				t.Errorf("InStrArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
