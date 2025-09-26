package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/images"
	"github.com/stretchr/testify/require"
)

func TestFit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	i := &imageResource{
		Image: &images.Image{},
	}

	// When
	_, err := i.Fit("")

	// Then
	require.NoError(t, err)
}
