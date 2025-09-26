package http

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientIPV2_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	err := clientIPV2(tree, "192.168.1.1")
	require.NoError(t, err)
	require.NotNil(t, tree.matcher)
}
