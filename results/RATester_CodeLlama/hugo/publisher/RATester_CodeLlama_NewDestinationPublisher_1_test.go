package publisher

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/media"
	"github.com/gohugoio/hugo/output"
	"github.com/gohugoio/hugo/resources"
	"github.com/stretchr/testify/require"
)

func TestNewDestinationPublisher_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	rs := &resources.Spec{}
	outputFormats := output.Formats{}
	mediaTypes := media.Types{}

	// When
	pub, err := NewDestinationPublisher(rs, outputFormats, mediaTypes)

	// Then
	require.NoError(t, err)
	require.NotNil(t, pub)
}
