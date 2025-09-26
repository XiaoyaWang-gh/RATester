package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestGetOrCreateBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Cache{
		Fs: afero.NewMemMapFs(),
	}

	id := "test"
	create := func() ([]byte, error) {
		return []byte("test"), nil
	}

	info, b, err := c.GetOrCreateBytes(id, create)
	require.NoError(t, err)
	require.Equal(t, "test", string(b))
	require.Equal(t, "test", info.Name)
}
