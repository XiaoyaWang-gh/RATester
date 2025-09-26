package client

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestAddFiles_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given:
	r := &Request{}
	files := []*File{
		{
			name:      "file1",
			fieldName: "file1",
			path:      "file1",
			reader:    nil,
		},
		{
			name:      "file2",
			fieldName: "file2",
			path:      "file2",
			reader:    nil,
		},
	}
	// when:
	r.AddFiles(files...)
	// then:
	assert.Equal(t, files, r.files)
}
