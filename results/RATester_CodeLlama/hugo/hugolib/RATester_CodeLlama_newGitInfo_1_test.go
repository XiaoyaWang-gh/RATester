package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
	"github.com/stretchr/testify/require"
)

func TestNewGitInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	d := &deps.Deps{}

	// when
	_, err := newGitInfo(d)

	// then
	require.NoError(t, err)
}
