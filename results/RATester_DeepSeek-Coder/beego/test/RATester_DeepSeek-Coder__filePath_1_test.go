package test

import (
	"fmt"
	"testing"
)

func Test_filePath_1(t *testing.T) {
	type args struct {
		dir  string
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				dir:  "/home/user/documents",
				name: "test/file.txt",
			},
			want: "/home/user/documents/test/file.txt",
		},
		{
			name: "Test case 2",
			args: args{
				dir:  "/home/user/documents/",
				name: "test/file.txt",
			},
			want: "/home/user/documents/test/file.txt",
		},
		{
			name: "Test case 3",
			args: args{
				dir:  "/home/user/documents",
				name: "test\\file.txt",
			},
			want: "/home/user/documents/test/file.txt",
		},
		{
			name: "Test case 4",
			args: args{
				dir:  "/home/user/documents/",
				name: "test\\file.txt",
			},
			want: "/home/user/documents/test/file.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := _filePath(tt.args.dir, tt.args.name); got != tt.want {
				t.Errorf("_filePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
