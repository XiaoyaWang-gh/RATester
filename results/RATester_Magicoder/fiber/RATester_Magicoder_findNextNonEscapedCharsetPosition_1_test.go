package fiber

import (
	"fmt"
	"testing"
)

func TestfindNextNonEscapedCharsetPosition_1(t *testing.T) {
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
			name: "Test case 1",
			args: args{
				search:  "\\a",
				charset: []byte{'a'},
			},
			want: -1,
		},
		{
			name: "Test case 2",
			args: args{
				search:  "a\\a",
				charset: []byte{'a'},
			},
			want: 1,
		},
		{
			name: "Test case 3",
			args: args{
				search:  "\\\\a",
				charset: []byte{'a'},
			},
			want: -1,
		},
		{
			name: "Test case 4",
			args: args{
				search:  "a\\\\a",
				charset: []byte{'a'},
			},
			want: 1,
		},
		{
			name: "Test case 5",
			args: args{
				search:  "\\\\\\a",
				charset: []byte{'a'},
			},
			want: -1,
		},
		{
			name: "Test case 6",
			args: args{
				search:  "a\\\\\\a",
				charset: []byte{'a'},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := findNextNonEscapedCharsetPosition(tt.args.search, tt.args.charset); got != tt.want {
				t.Errorf("findNextNonEscapedCharsetPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
