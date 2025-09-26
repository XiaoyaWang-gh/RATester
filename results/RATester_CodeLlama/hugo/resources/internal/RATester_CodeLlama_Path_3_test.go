package internal

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestPath_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := ResourcePaths{
		Dir:             "dir",
		File:            "file",
		BaseDirTarget:   "base",
		BaseDirLink:     "link",
		TargetBasePaths: []string{"base1", "base2"},
	}

	assert.Equal(t, "dir/file", d.Path())
}
