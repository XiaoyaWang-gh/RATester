package gin

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDir_1(t *testing.T) {
	type args struct {
		root          string
		listDirectory bool
	}
	tests := []struct {
		name string
		args args
		want http.FileSystem
	}{
		{
			name: "Test Case 1",
			args: args{
				root:          "/path/to/directory",
				listDirectory: true,
			},
			want: http.Dir("/path/to/directory"),
		},
		{
			name: "Test Case 2",
			args: args{
				root:          "/path/to/directory",
				listDirectory: false,
			},
			want: &OnlyFilesFS{FileSystem: http.Dir("/path/to/directory")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Dir(tt.args.root, tt.args.listDirectory); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dir() = %v, want %v", got, tt.want)
			}
		})
	}
}
