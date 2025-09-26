package common

import (
	"fmt"
	"testing"
)

func TestGetverifyval_1(t *testing.T) {
	type args struct {
		vkey string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{vkey: "test1"},
			want: "e10adc3949ba59abbe56e057f20f883e",
		},
		{
			name: "Test Case 2",
			args: args{vkey: "test2"},
			want: "e10adc3949ba59abbe56e057f20f883e",
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

			if got := Getverifyval(tt.args.vkey); got != tt.want {
				t.Errorf("Getverifyval() = %v, want %v", got, tt.want)
			}
		})
	}
}
