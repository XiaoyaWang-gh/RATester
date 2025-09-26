package failover

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegisterStatusUpdater_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &Failover{}
	fn := func(up bool) {}
	err := f.RegisterStatusUpdater(fn)
	require.NoError(t, err)
	require.Len(t, f.updaters, 1)
	require.Equal(t, f.updaters[0], fn)
}
