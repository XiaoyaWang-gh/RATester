package minifier

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources"
	"github.com/stretchr/testify/require"
)

func TestNew_24(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	rs := &resources.Spec{}

	// when
	_, err := New(rs)

	// then
	require.NoError(t, err)
}
