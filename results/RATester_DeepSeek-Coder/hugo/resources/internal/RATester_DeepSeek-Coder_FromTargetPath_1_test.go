package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFromTargetPath_1(t *testing.T) {
	type args struct {
		targetPath string
	}
	tests := []struct {
		name string
		d    ResourcePaths
		args args
		want ResourcePaths
	}{
		{
			name: "Test case 1",
			d:    ResourcePaths{},
			args: args{
				targetPath: "/path/to/file",
			},
			want: ResourcePaths{
				Dir:             "/path/to",
				BaseDirTarget:   "",
				BaseDirLink:     "",
				TargetBasePaths: []string{},
				File:            "file",
			},
		},
		{
			name: "Test case 2",
			d:    ResourcePaths{},
			args: args{
				targetPath: "/path/to/dir/",
			},
			want: ResourcePaths{
				Dir:             "/path/to/dir",
				BaseDirTarget:   "",
				BaseDirLink:     "",
				TargetBasePaths: []string{},
				File:            "",
			},
		},
		{
			name: "Test case 3",
			d:    ResourcePaths{},
			args: args{
				targetPath: "/path/to/dir/file.txt",
			},
			want: ResourcePaths{
				Dir:             "/path/to/dir",
				BaseDirTarget:   "",
				BaseDirLink:     "",
				TargetBasePaths: []string{},
				File:            "file.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := tt.d
			if got := d.FromTargetPath(tt.args.targetPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromTargetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
