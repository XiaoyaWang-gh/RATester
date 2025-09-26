package source

import (
	"fmt"
	"testing"
)

func TestpathToDir_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{s: "/test/path"},
			want: "/test/path/",
		},
		{
			name: "Test Case 2",
			args: args{s: ""},
			want: "",
		},
		{
			name: "Test Case 3",
			args: args{s: "/"},
			want: "/",
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
				t.Errorf("pathToDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
