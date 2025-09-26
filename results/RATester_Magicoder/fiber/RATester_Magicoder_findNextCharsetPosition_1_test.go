package fiber

import (
	"fmt"
	"testing"
)

func TestfindNextCharsetPosition_1(t *testing.T) {
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
				charset: []byte{'l', 'o'},
			},
			want: 3,
		},
		{
			name: "Test Case 2",
			args: args{
				search:  "hello world",
				charset: []byte{'h', 'w'},
			},
			want: 0,
		},
		{
			name: "Test Case 3",
			args: args{
				search:  "hello world",
				charset: []byte{'z'},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := findNextCharsetPosition(tt.args.search, tt.args.charset); got != tt.want {
				t.Errorf("findNextCharsetPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
