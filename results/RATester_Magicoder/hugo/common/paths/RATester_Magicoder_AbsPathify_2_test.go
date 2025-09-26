package paths

import (
	"fmt"
	"testing"
)

func TestAbsPathify_2(t *testing.T) {
	type args struct {
		workingDir string
		inPath     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Absolute path",
			args: args{
				workingDir: "/home/user",
				inPath:     "/usr/local/bin",
			},
			want: "/usr/local/bin",
		},
		{
			name: "Relative path",
			args: args{
				workingDir: "/home/user",
				inPath:     "bin",
			},
			want: "/home/user/bin",
		},
		{
			name: "Absolute path with trailing slash",
			args: args{
				workingDir: "/home/user/",
				inPath:     "/usr/local/bin",
			},
			want: "/usr/local/bin",
		},
		{
			name: "Relative path with trailing slash",
			args: args{
				workingDir: "/home/user/",
				inPath:     "bin/",
			},
			want: "/home/user/bin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := AbsPathify(tt.args.workingDir, tt.args.inPath); got != tt.want {
				t.Errorf("AbsPathify() = %v, want %v", got, tt.want)
			}
		})
	}
}
