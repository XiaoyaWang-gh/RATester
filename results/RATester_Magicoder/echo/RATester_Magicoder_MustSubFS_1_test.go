package echo

import (
	"fmt"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestMustSubFS_1(t *testing.T) {
	type args struct {
		currentFs fs.FS
		fsRoot    string
	}
	tests := []struct {
		name string
		args args
		want fs.FS
	}{
		{
			name: "Test case 1",
			args: args{
				currentFs: fstest.MapFS{
					"file1.txt": &fstest.MapFile{Data: []byte("Hello, World!")},
				},
				fsRoot: "file1.txt",
			},
			want: fstest.MapFS{
				"file1.txt": &fstest.MapFile{Data: []byte("Hello, World!")},
			},
		},
		{
			name: "Test case 2",
			args: args{
				currentFs: fstest.MapFS{
					"file1.txt": &fstest.MapFile{Data: []byte("Hello, World!")},
				},
				fsRoot: "file2.txt",
			},
			want: fstest.MapFS{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := MustSubFS(tt.args.currentFs, tt.args.fsRoot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustSubFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
