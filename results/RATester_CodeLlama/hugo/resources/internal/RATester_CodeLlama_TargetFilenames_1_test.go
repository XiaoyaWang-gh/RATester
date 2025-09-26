package internal

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestTargetFilenames_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := ResourcePaths{
		Dir:  "a/b/c",
		File: "data.json",
	}

	assert.Equal(t, []string{"a/b/c/data.json"}, d.TargetFilenames())
}
