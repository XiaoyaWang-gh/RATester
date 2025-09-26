package tcp

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestAlpnV2_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	err := alpnV2(tree, "http/1.1")
	assert.NoError(t, err)
	assert.NotNil(t, tree.matcher)
	assert.Equal(t, "http/1.1", tree.matcher(ConnData{alpnProtos: []string{"http/1.1"}}))
	assert.Nil(t, tree.left)
	assert.Nil(t, tree.right)
}
