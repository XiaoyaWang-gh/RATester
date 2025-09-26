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
			name: "Test case 1",
			args: args{
				s:      "Hello, world",
				start:  0,
				length: 5,
			},
			want: "Hello",
		},
		{
			name: "Test case 2",
			args: args{
				s:      "Hello, world",
				start:  7,
				length: 5,
			},
			want: "world",
		},
		{
			name: "Test case 3",
			args: args{
				s:      "Hello, world",
				start:  -1,
				length: 5,
			},
			want: "Hello",
		},
		{
			name: "Test case 4",
			args: args{
				s:      "Hello, world",
				start:  7,
				length: 15,
			},
			want: "world",
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
