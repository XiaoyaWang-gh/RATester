package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/cache/filecache"
	"github.com/stretchr/testify/require"
)

func TestWriteMeta_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	key := "key"
	meta := transformedResourceMetadata{
		Target:     "target",
		MediaTypeV: "mediaType",
		MetaData:   map[string]any{"key": "value"},
	}
	c := &ResourceCache{
		fileCache: &filecache.Cache{},
	}

	// When
	_, _, err := c.writeMeta(key, meta)

	// Then
	require.NoError(t, err)
}
