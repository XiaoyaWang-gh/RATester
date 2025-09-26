package tcp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientIP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	err := clientIP(tree, "192.168.1.1")
	require.NoError(t, err)
	require.NotNil(t, tree.matcher)
	require.Nil(t, tree.left)
	require.Nil(t, tree.right)
}
