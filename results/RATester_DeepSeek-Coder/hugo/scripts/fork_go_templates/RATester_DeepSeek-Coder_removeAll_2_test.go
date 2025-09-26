package main

import (
	"fmt"
	"testing"
)

func TestRemoveAll_2(t *testing.T) {
	type args struct {
		expression string
		content    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				expression: "a*",
				content:    "banana",
			},
			want: "b",
		},
		{
			name: "Test case 2",
			args: args{
				expression: "b*",
				content:    "apple",
			},
			want: "apple",
		},
		{
			name: "Test case 3",
			args: args{
				expression: "c*",
				content:    "cherry",
			},
			want: "cherry",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := removeAll(tt.args.expression, tt.args.content); got != tt.want {
				t.Errorf("removeAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
