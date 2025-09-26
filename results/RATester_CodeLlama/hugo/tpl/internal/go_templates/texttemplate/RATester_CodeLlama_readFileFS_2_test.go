package template

import (
	"fmt"
	"testing"
	"testing/fstest"
)

func TestReadFileFS_2(t *testing.T) {
	fsys := fstest.MapFS{
		"file1": &fstest.MapFile{Data: []byte("file1")},
		"file2": &fstest.MapFile{Data: []byte("file2")},
	}
	readFileFS := readFileFS(fsys)
	tests := []struct {
		name string
		file string
		want string
	}{
		{
			name: "file1",
			file: "file1",
			want: "file1",
		},
		{
			name: "file2",
			file: "file2",
			want: "file2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			name, b, err := readFileFS(tt.file)
			if err != nil {
				t.Fatal(err)
			}
			if name != tt.want {
				t.Errorf("readFileFS() name = %v, want %v", name, tt.want)
			}
			if string(b) != tt.want {
				t.Errorf("readFileFS() b = %v, want %v", string(b), tt.want)
			}
		})
	}
}
