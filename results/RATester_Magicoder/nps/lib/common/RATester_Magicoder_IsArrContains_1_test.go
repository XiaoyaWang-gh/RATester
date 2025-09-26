package common

import (
	"fmt"
	"testing"
)

func TestIsArrContains_1(t *testing.T) {
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
				arr: []string{"apple", "banana", "cherry"},
				val: "banana",
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				arr: []string{"apple", "banana", "cherry"},
				val: "orange",
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				arr: []string{},
				val: "banana",
			},
			want: false,
		},
		{
			name: "Test case 4",
			args: args{
				arr: nil,
				val: "banana",
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

			if got := IsArrContains(tt.args.arr, tt.args.val); got != tt.want {
				t.Errorf("IsArrContains() = %v, want %v", got, tt.want)
			}
		})
	}
}
