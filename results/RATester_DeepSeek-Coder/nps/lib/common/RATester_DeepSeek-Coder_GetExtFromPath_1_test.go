package common

import (
	"fmt"
	"testing"
)

func TestGetExtFromPath_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{path: "test.txt"},
			want: "txt",
		},
		{
			name: "Test case 2",
			args: args{path: "test.go"},
			want: "go",
		},
		{
			name: "Test case 3",
			args: args{path: "test.docx"},
			want: "docx",
		},
		{
			name: "Test case 4",
			args: args{path: "test"},
			want: "",
		},
		{
			name: "Test case 5",
			args: args{path: "test."},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetExtFromPath(tt.args.path); got != tt.want {
				t.Errorf("GetExtFromPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
