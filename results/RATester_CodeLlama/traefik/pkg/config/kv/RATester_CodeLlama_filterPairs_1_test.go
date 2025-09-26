package kv

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/kvtools/valkeyrie/store"
)

func TestFilterPairs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pairs := []*store.KVPair{
		{Key: "a/1", Value: []byte("1")},
		{Key: "a/2", Value: []byte("2")},
		{Key: "a/3", Value: []byte("3")},
		{Key: "b/1", Value: []byte("1")},
		{Key: "b/2", Value: []byte("2")},
		{Key: "b/3", Value: []byte("3")},
		{Key: "c/1", Value: []byte("1")},
		{Key: "c/2", Value: []byte("2")},
		{Key: "c/3", Value: []byte("3")},
	}

	filters := []string{"a", "b"}

	expected := []*store.KVPair{
		{Key: "a/1", Value: []byte("1")},
		{Key: "a/2", Value: []byte("2")},
		{Key: "a/3", Value: []byte("3")},
		{Key: "b/1", Value: []byte("1")},
		{Key: "b/2", Value: []byte("2")},
		{Key: "b/3", Value: []byte("3")},
	}

	actual := filterPairs(pairs, filters)

	assert.Equal(t, expected, actual)
}
