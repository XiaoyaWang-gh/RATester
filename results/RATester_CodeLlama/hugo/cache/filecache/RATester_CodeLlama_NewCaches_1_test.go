package filecache

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/helpers"
	"github.com/stretchr/testify/require"
)

func TestNewCaches_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	p := &helpers.PathSpec{}

	// When
	caches, err := NewCaches(p)

	// Then
	require.NoError(t, err)
	require.NotNil(t, caches)
}
