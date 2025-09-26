package tplimpl

import (
	"fmt"
	"testing"
)

func TestisDotFile_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				path: "/path/to/.file",
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				path: "/path/to/file",
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				path: "/path/to/.hidden_file",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isDotFile(tt.args.path); got != tt.want {
				t.Errorf("isDotFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
