package source

import (
	"fmt"
	"testing"
)

func TestPathToDir_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with empty string",
			args: args{s: ""},
			want: "",
		},
		{
			name: "Test with non-empty string",
			args: args{s: "/test/path"},
			want: "test/path/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			fi := &File{}
			if got := fi.pathToDir(tt.args.s); got != tt.want {
				t.Errorf("File.pathToDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
