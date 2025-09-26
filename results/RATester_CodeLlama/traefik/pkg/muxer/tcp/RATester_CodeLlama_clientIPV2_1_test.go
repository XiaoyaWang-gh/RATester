package tcp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientIPV2_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	err := clientIPV2(tree, "192.168.1.1")
	require.NoError(t, err)
	require.NotNil(t, tree.matcher)
	require.Nil(t, tree.left)
	require.Nil(t, tree.right)

	meta := ConnData{
		remoteIP: "192.168.1.1",
	}
	require.True(t, tree.matcher(meta))

	meta = ConnData{
		remoteIP: "192.168.1.2",
	}
	require.False(t, tree.matcher(meta))
}
