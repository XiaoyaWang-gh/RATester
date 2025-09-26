package plugins

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildProvider_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := Builder{}
	pName := "pName"
	config := map[string]interface{}{}

	p, err := b.BuildProvider(pName, config)
	require.NoError(t, err)
	require.NotNil(t, p)
}
