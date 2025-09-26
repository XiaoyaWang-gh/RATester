package scss

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeOptions_3(t *testing.T) {
	t.Parallel()

	t.Run("nil map", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		opts, err := DecodeOptions(nil)
		require.NoError(t, err)
		require.Empty(t, opts)
	})

	t.Run("empty map", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		opts, err := DecodeOptions(map[string]any{})
		require.NoError(t, err)
		require.Empty(t, opts)
	})

	t.Run("valid map", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		m := map[string]any{
			"targetPath": "foo",
			"includePaths": []any{
				"bar",
				"baz",
			},
			"outputStyle":     "compressed",
			"precision":       10,
			"enableSourceMap": true,
			"vars": map[string]any{
				"foo": "bar",
			},
		}

		opts, err := DecodeOptions(m)
		require.NoError(t, err)
		require.Equal(t, "foo", opts.TargetPath)
		require.Equal(t, []string{"bar", "baz"}, opts.IncludePaths)
		require.Equal(t, "compressed", opts.OutputStyle)
		require.Equal(t, 10, opts.Precision)
		require.True(t, opts.EnableSourceMap)
		require.Equal(t, map[string]any{"foo": "bar"}, opts.Vars)
	})
}
