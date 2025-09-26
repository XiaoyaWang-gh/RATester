package tpl

import (
	"fmt"
	"testing"
)

func TestextractBaseOf_1(t *testing.T) {
	type args struct {
		err string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				err: "base of /path/to/file.txt: line 1: syntax error",
			},
			want: "/path/to/file.txt",
		},
		{
			name: "Test case 2",
			args: args{
				err: "base of /path/to/file2.txt: line 2: syntax error",
			},
			want: "/path/to/file2.txt",
		},
		{
			name: "Test case 3",
			args: args{
				err: "base of /path/to/file3.txt: line 3: syntax error",
			},
			want: "/path/to/file3.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := extractBaseOf(tt.args.err); got != tt.want {
				t.Errorf("extractBaseOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
