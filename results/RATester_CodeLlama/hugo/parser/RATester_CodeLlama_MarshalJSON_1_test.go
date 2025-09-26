package parser

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalJSON_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := ReplacingJSONMarshaller{
		Value: map[string]any{
			"foo": "bar",
			"bar": "baz",
		},
		KeysToLower: true,
		OmitEmpty:   true,
	}

	// when
	result, err := c.MarshalJSON()

	// then
	require.NoError(t, err)
	assert.Equal(t, `{"foo":"bar","bar":"baz"}`, string(result))
}
