package file

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetDefaults_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	p.SetDefaults()
	require.True(t, p.Watch)
	require.Equal(t, "", p.Filename)
}
