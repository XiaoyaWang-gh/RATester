package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFileByPath_1(t *testing.T) {
	r := &Request{
		files: []*File{
			{path: "/path1"},
			{path: "/path2"},
		},
	}

	tests := []struct {
		name string
		path string
		want *File
	}{
		{
			name: "File exists",
			path: "/path1",
			want: &File{path: "/path1"},
		},
		{
			name: "File does not exist",
			path: "/path3",
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := r.FileByPath(tt.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileByPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
