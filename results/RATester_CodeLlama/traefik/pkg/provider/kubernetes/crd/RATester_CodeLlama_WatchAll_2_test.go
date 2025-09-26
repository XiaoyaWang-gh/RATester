package crd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWatchAll_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &clientWrapper{}
	namespaces := []string{"ns1", "ns2"}
	stopCh := make(chan struct{})

	// when
	eventCh, err := c.WatchAll(namespaces, stopCh)

	// then
	require.NoError(t, err)
	require.NotNil(t, eventCh)
}
