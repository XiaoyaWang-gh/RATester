package codegen

import (
	"fmt"
	"testing"
)

func TestfirstToLower_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				name: "Hello",
			},
			want: "hello",
		},
		{
			name: "Test Case 2",
			args: args{
				name: "WORLD",
			},
			want: "wORLD",
		},
		{
			name: "Test Case 3",
			args: args{
				name: "12345",
			},
			want: "12345",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := firstToLower(tt.args.name); got != tt.want {
				t.Errorf("firstToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}
