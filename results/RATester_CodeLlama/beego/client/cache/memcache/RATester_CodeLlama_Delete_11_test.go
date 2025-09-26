package memcache

import (
	"context"
	"fmt"
	"testing"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/zeebo/assert"
)

func TestDelete_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	ctx := context.Background()
	key := "key"
	value := "value"
	rc := &Cache{
		conn:     memcache.New("localhost:11211"),
		conninfo: []string{"localhost:11211"},
	}
	rc.Put(ctx, key, value, 0)

	// when
	err := rc.Delete(ctx, key)

	// then
	assert.NoError(t, err)
	_, err = rc.Get(ctx, key)
	assert.Error(t, err)
}
