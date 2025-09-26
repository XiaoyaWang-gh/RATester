package collector

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestCollect_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	staticConfiguration := &static.Configuration{}

	// when
	err := Collect(staticConfiguration)

	// then
	require.NoError(t, err)
}
