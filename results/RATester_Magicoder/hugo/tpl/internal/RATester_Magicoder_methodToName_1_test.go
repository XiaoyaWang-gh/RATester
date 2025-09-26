package internal

import (
	"fmt"
	"testing"
)

func TestmethodToName_1(t *testing.T) {
	type args struct {
		m any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				m: func() {},
			},
			want: "func1",
		},
		{
			name: "Test case 2",
			args: args{
				m: func() {},
			},
			want: "func2",
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

			if got := methodToName(tt.args.m); got != tt.want {
				t.Errorf("methodToName() = %v, want %v", got, tt.want)
			}
		})
	}
}
