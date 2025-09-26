package helpers

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestOpenFileForWriting_1(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// given
		fs := afero.NewMemMapFs()
		filename := "test.txt"

		// when
		f, err := OpenFileForWriting(fs, filename)

		// then
		require.NoError(t, err)
		require.NotNil(t, f)
	})

	t.Run("failure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// given
		fs := afero.NewMemMapFs()
		filename := "test.txt"

		// when
		f, err := OpenFileForWriting(fs, filename)

		// then
		require.NoError(t, err)
		require.NotNil(t, f)
	})
}
