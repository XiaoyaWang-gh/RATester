package paths

import (
	"fmt"
	"testing"
)

func TestIsSameFilePath_1(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				s1: "/path/to/file",
				s2: "/path/to/file",
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				s1: "/path/to/file/",
				s2: "/path/to/file",
			},
			want: true,
		},
		{
			name: "Test case 3",
			args: args{
				s1: "/path/to/file",
				s2: "/path/to/file/",
			},
			want: true,
		},
		{
			name: "Test case 4",
			args: args{
				s1: "/path/to/file/",
				s2: "/path/to/file/",
			},
			want: true,
		},
		{
			name: "Test case 5",
			args: args{
				s1: "/path/to/file1",
				s2: "/path/to/file2",
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

			if got := IsSameFilePath(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("IsSameFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
