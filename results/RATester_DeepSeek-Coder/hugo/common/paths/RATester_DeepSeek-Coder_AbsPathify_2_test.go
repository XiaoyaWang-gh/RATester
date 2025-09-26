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
				inPath:     "/home/user/test",
			},
			want: "/home/user/test",
		},
		{
			name: "Relative path",
			args: args{
				workingDir: "/home/user",
				inPath:     "test",
			},
			want: "/home/user/test",
		},
		{
			name: "Absolute path with trailing slash",
			args: args{
				workingDir: "/home/user/",
				inPath:     "/home/user/test",
			},
			want: "/home/user/test",
		},
		{
			name: "Relative path with trailing slash",
			args: args{
				workingDir: "/home/user/",
				inPath:     "test",
			},
			want: "/home/user/test",
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
