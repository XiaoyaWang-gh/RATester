package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTargetPaths_1(t *testing.T) {
	tests := []struct {
		name string
		d    ResourcePaths
		want []string
	}{
		{
			name: "single target base path",
			d: ResourcePaths{
				TargetBasePaths: []string{"/path/to/target"},
				Dir:             "dir",
				File:            "file.txt",
			},
			want: []string{"/path/to/target/dir/file.txt"},
		},
		{
			name: "multiple target base paths",
			d: ResourcePaths{
				TargetBasePaths: []string{"/path/to/target1", "/path/to/target2"},
				Dir:             "dir",
				File:            "file.txt",
			},
			want: []string{"/path/to/target1/dir/file.txt", "/path/to/target2/dir/file.txt"},
		},
		{
			name: "no target base paths",
			d: ResourcePaths{
				Dir:  "dir",
				File: "file.txt",
			},
			want: []string{"/dir/file.txt"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.d.TargetPaths(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourcePaths.TargetPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
