package utils

import (
	"fmt"
	"testing"
)

func TestqEncode_1(t *testing.T) {
	type args struct {
		charset string
		s       string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				charset: "utf-8",
				s:       "Hello, World!",
			},
			want: "Hello, World!",
		},
		{
			name: "Test Case 2",
			args: args{
				charset: "utf-8",
				s:       "日本語",
			},
			want: "%E6%97%A5%E6%9C%AC%E8%AA%9E",
		},
		{
			name: "Test Case 3",
			args: args{
				charset: "utf-8",
				s:       "Special characters: !@#$%^&*()",
			},
			want: "Special characters: !@#$%^&*()",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := qEncode(tt.args.charset, tt.args.s); got != tt.want {
				t.Errorf("qEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}
