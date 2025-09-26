package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTargetPaths_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := ResourcePaths{
		Dir:             "dir",
		BaseDirTarget:   "base",
		BaseDirLink:     "link",
		TargetBasePaths: []string{"base1", "base2"},
		File:            "file",
	}

	if got, want := d.TargetPaths(), []string{"base/dir/file", "link/dir/file", "base1/dir/file", "base2/dir/file"}; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
