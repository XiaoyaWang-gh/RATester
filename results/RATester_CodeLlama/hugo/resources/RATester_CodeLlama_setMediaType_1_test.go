package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/media"
)

func TestSetMediaType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	l := &genericResource{}
	mediaType := media.Type{
		Type: "application/rss+xml",
	}

	// When
	l.setMediaType(mediaType)

	// Then
	assert.Equal(t, mediaType, l.MediaType())
}
