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

	ctx := context.Background()
	mc := &Cache{
		conn:     memcache.New("127.0.0.1:11211"),
		conninfo: []string{"127.0.0.1:11211"},
	}

	err := mc.ClearAll(ctx)
	if err != nil {
		t.Errorf("ClearAll() error = %v, wantErr %v", err, false)
		return
	}
}
