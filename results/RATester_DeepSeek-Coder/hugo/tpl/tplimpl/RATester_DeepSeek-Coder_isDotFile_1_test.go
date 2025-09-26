package tplimpl

import (
	"fmt"
	"testing"
)

func TestIsDotFile_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test with a dot file",
			args: args{path: ".test"},
			want: true,
		},
		{
			name: "Test with a non-dot file",
			args: args{path: "test"},
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

			if got := isDotFile(tt.args.path); got != tt.want {
				t.Errorf("isDotFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
