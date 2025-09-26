package gateway

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListTLSRoutes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &clientWrapper{
		watchedNamespaces: []string{"ns1", "ns2"},
	}

	// when
	routes, err := c.ListTLSRoutes()

	// then
	require.NoError(t, err)
	require.Len(t, routes, 2)
}
