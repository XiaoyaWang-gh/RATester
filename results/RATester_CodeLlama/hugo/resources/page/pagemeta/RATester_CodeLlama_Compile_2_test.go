package pagemeta

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/common/paths"
	"github.com/gohugoio/hugo/media"
	"github.com/stretchr/testify/require"
)

func TestCompile_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	rc := &ResourceConfig{
		Path: "path/to/file.md",
	}
	basePath := "/base/path"
	pathParser := &paths.PathParser{}
	mediaTypes := media.Types{}

	// When
	err := rc.Compile(basePath, pathParser, mediaTypes)

	// Then
	require.NoError(t, err)
	assert.Equal(t, "/base/path/path/to/file.md", rc.Path)
}
