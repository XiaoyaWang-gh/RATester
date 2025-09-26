package internal

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestTargetPath_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := ResourcePaths{
		Dir:             "a/b",
		BaseDirTarget:   "c/d",
		BaseDirLink:     "e/f",
		TargetBasePaths: []string{"g/h", "i/j"},
		File:            "k.json",
	}

	assert.Equal(t, "c/d/a/b/k.json", d.TargetPath())
}
