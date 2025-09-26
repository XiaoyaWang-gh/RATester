package crd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewExternalClusterClientFromFile_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	file := "testdata/kubeconfig.yaml"

	// when
	client, err := newExternalClusterClientFromFile(file)

	// then
	require.NoError(t, err)
	require.NotNil(t, client)
}
