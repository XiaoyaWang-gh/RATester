package filecache

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestWriteCloser_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Cache{
		Fs: afero.NewMemMapFs(),
	}

	id := "foo"

	info, w, err := c.WriteCloser(id)
	require.NoError(t, err)
	require.NotNil(t, w)
	require.NotNil(t, info)
	require.Equal(t, id, info.Name)

	_, err = w.Write([]byte("hello"))
	require.NoError(t, err)

	err = w.Close()
	require.NoError(t, err)

	f, err := c.Fs.Open(id)
	require.NoError(t, err)
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	require.NoError(t, err)
	require.Equal(t, "hello", string(b))
}
