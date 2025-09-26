package openapi3

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
	"github.com/stretchr/testify/require"
)

func TestNew_58(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	deps := &deps.Deps{}
	ns := New(deps)
	require.NotNil(t, ns)
	require.NotNil(t, ns.cache)
	require.NotNil(t, ns.deps)
}
