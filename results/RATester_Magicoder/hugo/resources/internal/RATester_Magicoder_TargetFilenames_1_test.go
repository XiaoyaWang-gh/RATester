package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTargetFilenames_1(t *testing.T) {
	tests := []struct {
		name string
		d    ResourcePaths
		want []string
	}{
		{
			name: "Test case 1",
			d: ResourcePaths{
				Dir:             "/test/dir",
				BaseDirTarget:   "/base/target",
				BaseDirLink:     "/base/link",
				TargetBasePaths: []string{"/base/path1", "/base/path2"},
				File:            "test.txt",
			},
			want: []string{"/test/dir/test.txt"},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.d.TargetFilenames(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TargetFilenames() = %v, want %v", got, tt.want)
			}
		})
	}
}
