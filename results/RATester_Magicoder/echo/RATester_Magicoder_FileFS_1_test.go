package echo

import (
	"fmt"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestFileFS_1(t *testing.T) {
	type test struct {
		name       string
		file       string
		filesystem fs.FS
		wantErr    bool
	}

	tests := []test{
		{
			name: "Test case 1",
			file: "test.txt",
			filesystem: fstest.MapFS{
				"test.txt": &fstest.MapFile{Data: []byte("Hello, World!")},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			file: "nonexistent.txt",
			filesystem: fstest.MapFS{
				"test.txt": &fstest.MapFile{Data: []byte("Hello, World!")},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &context{}
			err := c.FileFS(tt.file, tt.filesystem)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileFS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
