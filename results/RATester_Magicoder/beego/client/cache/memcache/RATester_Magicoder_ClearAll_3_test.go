package memcache

import (
	"context"
	"fmt"
	"testing"

	"github.com/bradfitz/gomemcache/memcache"
)

func TestClearAll_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mc := &Cache{
		conn: &memcache.Client{},
	}

	ctx := context.Background()
	err := mc.ClearAll(ctx)
	if err != nil {
		t.Errorf("ClearAll() error = %v, wantErr %v", err, false)
		return
	}
}
