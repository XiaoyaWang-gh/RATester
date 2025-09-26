package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTargetFilenames_1(t *testing.T) {
	type fields struct {
		Dir             string
		BaseDirTarget   string
		BaseDirLink     string
		TargetBasePaths []string
		File            string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "Test 1",
			fields: fields{
				Dir:             "/dir",
				BaseDirTarget:   "/base",
				BaseDirLink:     "/link",
				TargetBasePaths: []string{"/path1", "/path2"},
				File:            "file.txt",
			},
			want: []string{"/dir/file.txt", "/base/file.txt", "/link/file.txt", "/path1/file.txt", "/path2/file.txt"},
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := ResourcePaths{
				Dir:             tt.fields.Dir,
				BaseDirTarget:   tt.fields.BaseDirTarget,
				BaseDirLink:     tt.fields.BaseDirLink,
				TargetBasePaths: tt.fields.TargetBasePaths,
				File:            tt.fields.File,
			}
			if got := d.TargetFilenames(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TargetFilenames() = %v, want %v", got, tt.want)
			}
		})
	}
}
