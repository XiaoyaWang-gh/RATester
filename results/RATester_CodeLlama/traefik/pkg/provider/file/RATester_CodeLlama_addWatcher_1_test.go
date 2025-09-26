package file

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/safe"
)

func TestAddWatcher_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	p := &Provider{}
	pool := &safe.Pool{}
	items := []string{"test"}
	configurationChan := make(chan dynamic.Message)
	callback := func(chan<- dynamic.Message) error {
		return nil
	}

	// when
	err := p.addWatcher(pool, items, configurationChan, callback)

	// then
	require.NoError(t, err)
}
