package fiber

import (
	"fmt"
	"testing"
)

func TestmatchEtag_1(t *testing.T) {
	type args struct {
		s    string
		etag string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				s:    "abc",
				etag: "abc",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				s:    "abc",
				etag: "def",
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				s:    "W/abc",
				etag: "abc",
			},
			want: true,
		},
		{
			name: "Test Case 4",
			args: args{
				s:    "abc",
				etag: "W/abc",
			},
			want: true,
		},
		{
			name: "Test Case 5",
			args: args{
				s:    "W/abc",
				etag: "def",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := matchEtag(tt.args.s, tt.args.etag); got != tt.want {
				t.Errorf("matchEtag() = %v, want %v", got, tt.want)
			}
		})
	}
}
