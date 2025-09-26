package filesystems

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/spf13/afero"
)

func TestResolveJSConfigFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	var (
		b    *BaseFs
		name string
	)

	// CONTEXT_0
	b = &BaseFs{}

	// CONTEXT_1
	b.Assets = &SourceFilesystem{}
	b.Work = afero.NewMemMapFs()

	// CONTEXT_2
	name = "test.js"

	// TEST CASE
	assert.Equal(t, "", b.ResolveJSConfigFile(name))
}
