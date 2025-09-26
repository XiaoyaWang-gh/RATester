package kv

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/kvtools/valkeyrie/store"
	"github.com/stretchr/testify/require"
)

func TestDecode_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	pairs := []*store.KVPair{
		{
			Key:       "key1",
			Value:     []byte("value1"),
			LastIndex: 1,
		},
		{
			Key:       "key2",
			Value:     []byte("value2"),
			LastIndex: 2,
		},
	}
	element := &struct {
		Key1 string `kv:"key1"`
		Key2 string `kv:"key2"`
	}{}
	rootName := "root"

	// when
	err := Decode(pairs, element, rootName)

	// then
	require.NoError(t, err)
	assert.Equal(t, "value1", element.Key1)
	assert.Equal(t, "value2", element.Key2)
}
