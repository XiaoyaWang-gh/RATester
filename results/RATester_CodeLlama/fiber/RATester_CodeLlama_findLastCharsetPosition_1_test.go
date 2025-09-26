package fiber

import (
	"fmt"
	"testing"
)

func TestFindLastCharsetPosition_1(t *testing.T) {
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
				search:  "test",
				charset: []byte("abc"),
			},
			want: 3,
		},
		{
			name: "test case 2",
			args: args{
				search:  "test",
				charset: []byte("abcd"),
			},
			want: 3,
		},
		{
			name: "test case 3",
			args: args{
				search:  "test",
				charset: []byte("abcd"),
			},
			want: 3,
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
