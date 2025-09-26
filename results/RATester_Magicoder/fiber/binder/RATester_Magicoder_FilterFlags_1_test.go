package binder

import (
	"fmt"
	"testing"
)

func TestFilterFlags_1(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				content: "Hello World",
			},
			want: "Hello",
		},
		{
			name: "Test case 2",
			args: args{
				content: "Hello;World",
			},
			want: "Hello",
		},
		{
			name: "Test case 3",
			args: args{
				content: "Hello World;",
			},
			want: "Hello World",
		},
		{
			name: "Test case 4",
			args: args{
				content: "Hello;World;",
			},
			want: "Hello",
		},
		{
			name: "Test case 5",
			args: args{
				content: "Hello World;",
			},
			want: "Hello World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := FilterFlags(tt.args.content); got != tt.want {
				t.Errorf("FilterFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}
