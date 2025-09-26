package web

import (
	"fmt"
	"testing"
)

func TestSubstr_1(t *testing.T) {
	type args struct {
		s      string
		start  int
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				s:      "Hello, World!",
				start:  7,
				length: 5,
			},
			want: "World",
		},
		{
			name: "Test Case 2",
			args: args{
				s:      "Hello, World!",
				start:  0,
				length: 5,
			},
			want: "Hello",
		},
		{
			name: "Test Case 3",
			args: args{
				s:      "Hello, World!",
				start:  -1,
				length: 5,
			},
			want: "Hello",
		},
		{
			name: "Test Case 4",
			args: args{
				s:      "Hello, World!",
				start:  10,
				length: 5,
			},
			want: "orld!",
		},
		{
			name: "Test Case 5",
			args: args{
				s:      "Hello, World!",
				start:  10,
				length: 10,
			},
			want: "orld!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Substr(tt.args.s, tt.args.start, tt.args.length); got != tt.want {
				t.Errorf("Substr() = %v, want %v", got, tt.want)
			}
		})
	}
}
