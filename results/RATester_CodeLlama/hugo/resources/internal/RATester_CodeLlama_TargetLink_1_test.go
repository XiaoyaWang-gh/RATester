package internal

import (
	"fmt"
	"testing"
)

func TestTargetLink_1(t *testing.T) {
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

	if got, want := d.TargetLink(), "link/dir/file"; got != want {
		t.Errorf("d.TargetLink() = %q, want %q", got, want)
	}
}
