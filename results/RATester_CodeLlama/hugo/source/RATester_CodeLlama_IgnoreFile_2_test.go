package source

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/spf13/afero"
)

func TestIgnoreFile_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var filename string
	var s *SourceSpec

	// CONTEXT_0
	s = &SourceSpec{}

	// CONTEXT_1
	s = &SourceSpec{
		SourceFs: afero.NewMemMapFs(),
	}

	// CONTEXT_2
	filename = "foo.txt"

	// CONTEXT_3
	s.SourceFs = afero.NewMemMapFs()

	// TEST CASE
	assert.True(t, s.IgnoreFile(filename))
}
