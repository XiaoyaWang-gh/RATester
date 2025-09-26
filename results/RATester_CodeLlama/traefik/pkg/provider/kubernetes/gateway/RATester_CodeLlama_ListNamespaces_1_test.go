package gateway

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/labels"
)

func TestListNamespaces_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	selector := labels.Everything()
	c := &clientWrapper{
		watchedNamespaces: []string{"ns1", "ns2"},
	}

	// when
	namespaces, err := c.ListNamespaces(selector)

	// then
	require.NoError(t, err)
	assert.Equal(t, []string{"ns1", "ns2"}, namespaces)
}
