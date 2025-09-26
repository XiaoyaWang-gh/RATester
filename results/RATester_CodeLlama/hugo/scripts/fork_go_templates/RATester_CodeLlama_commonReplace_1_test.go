package main

import (
	"fmt"
	"testing"
)

func TestCommonReplace_1(t *testing.T) {
	type args struct {
		name    string
		content string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				name:    "test1.go",
				content: "package template\n",
			},
			want: "package template\n",
		},
		{
			name: "test2",
			args: args{
				name:    "test2.go",
				content: "package template_test\n",
			},
			want: "package template_test\n",
		},
		{
			name: "test3",
			args: args{
				name:    "test3.go",
				content: "package parse\n",
			},
			want: "package parse\n",
		},
		{
			name: "test4",
			args: args{
				name:    "test4.go",
				content: "package parse_test\n",
			},
			want: "package parse_test\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := commonReplace(tt.args.name, tt.args.content); got != tt.want {
				t.Errorf("commonReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}
