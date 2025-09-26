package fiber

import (
	"fmt"
	"testing"
)

func TestfindLastCharsetPosition_1(t *testing.T) {
	type args struct {
		search  string
		charset []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				search:  "hello world",
				charset: []byte{'o', 'l'},
			},
			want: 8,
		},
		{
			name: "Test Case 2",
			args: args{
				search:  "hello world",
				charset: []byte{'z'},
			},
			want: -1,
		},
		{
			name: "Test Case 3",
			args: args{
				search:  "hello world",
				charset: []byte{'l', 'o'},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := findLastCharsetPosition(tt.args.search, tt.args.charset); got != tt.want {
				t.Errorf("findLastCharsetPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
