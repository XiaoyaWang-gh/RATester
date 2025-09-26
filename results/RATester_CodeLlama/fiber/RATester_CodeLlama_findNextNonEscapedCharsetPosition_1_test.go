package fiber

import (
	"fmt"
	"testing"
)

func TestFindNextNonEscapedCharsetPosition_1(t *testing.T) {
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
			name: "test case 1",
			args: args{
				search: "test",
				charset: []byte{
					't',
					'e',
					's',
					't',
				},
			},
			want: 4,
		},
		{
			name: "test case 2",
			args: args{
				search: "test",
				charset: []byte{
					't',
					'e',
					's',
					't',
				},
			},
			want: 4,
		},
		{
			name: "test case 3",
			args: args{
				search: "test",
				charset: []byte{
					't',
					'e',
					's',
					't',
				},
			},
			want: 4,
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
