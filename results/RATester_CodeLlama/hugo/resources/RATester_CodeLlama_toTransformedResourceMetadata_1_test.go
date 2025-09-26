package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/media"
)

func TestToTransformedResourceMetadata_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	u := &transformationUpdate{}
	u.mediaType = media.Type{Type: "text/html"}
	u.targetPath = "targetPath"
	u.data = map[string]any{"key": "value"}
	actual := u.toTransformedResourceMetadata()
	expected := transformedResourceMetadata{
		Target:     "targetPath",
		MediaTypeV: "text/html",
		MetaData:   map[string]any{"key": "value"},
	}
	assert.Equal(t, expected, actual)
}
